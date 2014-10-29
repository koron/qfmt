package qfmt

import (
	"io"
	"os"
)

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	f, err := New(format)
	if err != nil {
		return 0, err
	}
	return f.Format(to_w(w), a...)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

func Sprintf(format string, a ...interface{}) string {
	b := bb_get()
	Fprintf(b, format, a...)
	s := b.String()
	bb_put(b)
	return s
}
