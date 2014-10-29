package qfmt

import (
	"errors"
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

type vemitter func(w Writer, a interface{}) (n int, err error)

func (v vemitter) bind(idx int) emitter {
	return func(w Writer, a []interface{}) (n int, err error) {
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

func emit_integer(w Writer, a interface{}) (n int, err error) {
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

func emit_string(w Writer, a interface{}) (n int, err error) {
	switch v := a.(type) {
	case string:
		n, err = w.WriteString(v)
	case Stringer:
		n, err = w.WriteString(v.String())
	default:
		n, err = emit_badverb(w, "s", a)
	}
	return
}

func emit_int64(w Writer, v int64) (n int, err error) {
	s := strconv.FormatInt(v, 10)
	return w.WriteString(s)
}

func emit_uint64(w Writer, v uint64) (n int, err error) {
	s := strconv.FormatUint(v, 10)
	return w.WriteString(s)
}

// emit_value emit in %v format.
func emit_value(w Writer, a interface{}) (n int, err error) {
	// TODO: implement emit_value
	return
}

func emit_reflect_value(w Writer, v reflect.Value) (n int, err error) {
	// TODO: implement emit_reflect_value
	return
}

func emit_badverb(w Writer, v string, a interface{}) (n int, err error) {
	w.WriteString("%!")
	w.WriteString(v)
	w.WriteString("(")
	if a != nil {
		w.WriteString(reflect.TypeOf(a).String())
		w.WriteString("=")
		emit_value(w, a)
	} else if rv := reflect.ValueOf(a); rv.IsValid() {
		w.WriteString(rv.Type().String())
		w.WriteString("=")
		emit_reflect_value(w, rv)
	} else {
		w.Write(nilAngleBytes)
	}
	w.WriteString(")")
	return
}
