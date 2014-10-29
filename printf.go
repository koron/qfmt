package qfmt

import (
	"bytes"
	"io"
	"os"
	"sync"
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

var bbPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

func Sprintf(format string, a ...interface{}) string {
	b := buf_get()
	Fprintf(b, format, a...)
	s := string(*b)
	buf_put(b)
	return s
}
