package qfmt

import (
	"errors"
	"io"
	"strconv"
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
	// TODO: format integer.
	return 0, nil
}

func emit_string(w io.Writer, a interface{}) (n int, err error) {
	if s, ok := a.(string); ok {
		return w.Write([]byte(s))
	}
	if s, ok := a.(Stringer); ok {
		return w.Write([]byte(s.String()))
	}
	return 0, nil
}
