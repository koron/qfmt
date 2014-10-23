package qfmt

import (
	"errors"
	"io"
	"regexp"
	"strings"
)

type Formatter interface {
	Format(w io.Writer, a ...interface{}) (n int, err error)
	//Bind(a ...interface{}) (Formatter, error)
}

// NOTE: copied from fmt/print.go
type Stringer interface {
	String() string
}

type emitter func(w io.Writer, a []interface{}) (n int, err error)

type fmtr struct {
	format   string
	emitters []emitter
}

func New(format string) (Formatter, error) {
	emitters, err := toEmitters(format)
	if err != nil {
		return nil, err
	}
	f := fmtr{
		format:   format,
		emitters: emitters,
	}
	return &f, nil
}

func (f *fmtr) Format(w io.Writer, a ...interface{}) (n int, err error) {
	for _, e := range f.emitters {
		m, err := e(w, a)
		n += m
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

func countTokens(s string) int {
	return strings.Count(s, "%")
}

func toEmitters(format string) (emitters []emitter, err error) {
	emitters = make([]emitter, 0, countTokens(format)*2+1)
	idx := 0
	for i := 0; i < len(format); {
		e, t, n, err := toEmitter(format[i:], idx)
		if err != nil {
			return nil, err
		}
		if t == "" {
			return nil, errors.New("empty token: " + format[i:])
		}
		idx += n
		emitters = append(emitters, e)
		i += len(t)
	}
	return emitters, nil
}

var rx_const = regexp.MustCompile(`^(?:%%|[^%])+`)
var rx_verb = regexp.MustCompile(`^%([ds])`)

func toEmitter(s string, idx int) (e emitter, token string, nargs int, err error) {
	if m := rx_const.FindString(s); m != "" {
		return const_emitter(m), m, 0, nil
	} else if m := rx_verb.FindStringSubmatch(s); m != nil {
		ve, err := verb2emitter(m[1])
		if err != nil {
			return nil, "", 0, err
		}
		return ve.bind(idx), m[0], 1, nil
	} else {
		return nil, "", 0, errors.New("unknown format: " + s)
	}
}

func const_emitter(s string) emitter {
	return func(w io.Writer, _ []interface{}) (n int, err error) {
		return w.Write([]byte(s))
	}
}
