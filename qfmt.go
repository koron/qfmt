package qfmt

import (
	"io"
	"strings"
)

type Formatter interface {
	Format(w io.Writer, a ...interface{}) (n int, err error)
	//Bind(a ...interface{}) (Formatter, error)
}

type emitter func(w io.Writer, a ...interface{}) (n int, err error)

type fmt struct {
	format   string
	emitters []emitter
}

func New(format string) (Formatter, error) {
	emitters, err := toEmitters(format)
	if err != nil {
		return nil, err
	}
	f := fmt{
		format:   format,
		emitters: emitters,
	}
	return &f, nil
}

func (f *fmt) Format(w io.Writer, a ...interface{}) (n int, err error) {
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
	for i = 0; i < len(format); i += 1 {
	}
	// TODO:
	return emitters, nil
}
