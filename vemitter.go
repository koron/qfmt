package qfmt

import (
	"errors"
	"io"
	"reflect"
	"strconv"
)

// Some constants in the form of bytes, to avoid string overhead.
// Needlessly fastidious, I suppose.
var (
	commaSpaceBytes  = []byte(", ")
	nilAngleBytes    = []byte("<nil>")
	nilParenBytes    = []byte("(nil)")
	nilBytes         = []byte("nil")
	mapBytes         = []byte("map[")
	percentBangBytes = []byte("%!")
	missingBytes     = []byte("(MISSING)")
	badIndexBytes    = []byte("(BADINDEX)")
	panicBytes       = []byte("(PANIC=")
	extraBytes       = []byte("%!(EXTRA ")
	irparenBytes     = []byte("i)")
	bytesBytes       = []byte("[]byte{")
	badWidthBytes    = []byte("%!(BADWIDTH)")
	badPrecBytes     = []byte("%!(BADPREC)")
	noVerbBytes      = []byte("%!(NOVERB)")
)

type vemitter func(w io.Writer, a interface{}) (n int, err error)

func (v vemitter) bind(idx int) emitter {
	return func(w io.Writer, a []interface{}) (n int, err error) {
		if idx < 0 || idx >= len(a) {
			return 0, errors.New("out of range:" + strconv.Itoa(idx))
		}
		return v(w, a[idx])
	}
}

func verb2emitter(verb string) (vemitter, error) {
	switch verb {
	case "d":
		return emit_integer, nil
	case "s":
		return emit_string, nil

	// TODO: add emitter for each types

	default:
		return nil, errors.New("unknown verb:" + verb)
	}
}

func emit_integer(w io.Writer, a interface{}) (n int, err error) {
	switch f := a.(type) {
	case int:
		n, err = emit_int64(w, int64(f))
	case int8:
		n, err = emit_int64(w, int64(f))
	case int16:
		n, err = emit_int64(w, int64(f))
	case int32:
		n, err = emit_int64(w, int64(f))
	case int64:
		n, err = emit_int64(w, int64(f))
	case uint:
		n, err = emit_uint64(w, uint64(f))
	case uint8:
		n, err = emit_uint64(w, uint64(f))
	case uint16:
		n, err = emit_uint64(w, uint64(f))
	case uint32:
		n, err = emit_uint64(w, uint64(f))
	case uint64:
		n, err = emit_uint64(w, uint64(f))
	default:
		n, err = emit_badverb(w, "d", a)
	}
	return
}

func emit_string(w io.Writer, a interface{}) (n int, err error) {
	switch v := a.(type) {
	case string:
		n, err = io.WriteString(w, v)
	case Stringer:
		n, err = io.WriteString(w, v.String())
	default:
		n, err = emit_badverb(w, "s", a)
	}
	return
}

func emit_int64(w io.Writer, v int64) (n int, err error) {
	s := strconv.FormatInt(v, 10)
	return io.WriteString(w, s)
}

func emit_uint64(w io.Writer, v uint64) (n int, err error) {
	s := strconv.FormatUint(v, 10)
	return io.WriteString(w, s)
}

// emit_value emit in %v format.
func emit_value(w io.Writer, a interface{}) (n int, err error) {
	// TODO: implement emit_value
	return
}

func emit_reflect_value(w io.Writer, v reflect.Value) (n int, err error) {
	// TODO: implement emit_reflect_value
	return
}

func emit_badverb(w io.Writer, v string, a interface{}) (n int, err error) {
	io.WriteString(w, "%!")
	io.WriteString(w, v)
	io.WriteString(w, "(")
	if a != nil {
		io.WriteString(w, reflect.TypeOf(a).String())
		io.WriteString(w, "=")
		emit_value(w, a)
	} else if rv := reflect.ValueOf(a); rv.IsValid() {
		io.WriteString(w, rv.Type().String())
		io.WriteString(w, "=")
		emit_reflect_value(w, rv)
	} else {
		w.Write(nilAngleBytes)
	}
	io.WriteString(w, ")")
	return
}
