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

func type2emitter(typename string) (vemitter, error) {
	switch typename {
	case "d":
		return emit_integer, nil
	case "s":
		return emit_string, nil

	// TODO: add emitter for each types

	default:
		return nil, errors.New("unknown type:" + typename)
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
	if s, ok := a.(string); ok {
		return w.Write([]byte(s))
	}
	if s, ok := a.(Stringer); ok {
		return w.Write([]byte(s.String()))
	}
	return emit_badverb(w, "s", a)
}

func emit_int64(w io.Writer, v int64) (n int, err error) {
	s := strconv.FormatInt(v, 10)
	return w.Write([]byte(s))
}

func emit_uint64(w io.Writer, v uint64) (n int, err error) {
	s := strconv.FormatUint(v, 10)
	return w.Write([]byte(s))
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
	w.Write([]byte("%!"))
	w.Write([]byte(v))
	w.Write([]byte("("))
	if a != nil {
		w.Write([]byte(reflect.TypeOf(a).String()))
		w.Write([]byte("="))
		emit_value(w, a)
	} else if rv := reflect.ValueOf(a); rv.IsValid() {
		w.Write([]byte(rv.Type().String()))
		w.Write([]byte("="))
		emit_reflect_value(w, rv)
	} else {
		w.Write(nilAngleBytes)
	}
	w.Write([]byte(")"))
	return
}
