package qfmt

import (
	"bytes"
	"io"
	"os"
)

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	f, err := New(format)
	if err != nil {
		return 0, err
	}
	return f.Format(w, a...)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

func Sprintf(format string, a ...interface{}) string {
	b := new(bytes.Buffer)
	_, err := Fprintf(b, format, a...)
	if err != nil {
		return ""
	}
	return b.String()
}
