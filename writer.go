package qfmt

import "io"

type Writer interface {
	Write(p []byte) (n int, err error)
	WriteString(s string) (n int, err error)
}

func to_w(iow io.Writer) (w Writer) {
	switch w := iow.(type) {
	case Writer:
		return w
	default:
		return &wrapWrite{writer: iow}
	}
}

type wrapWrite struct {
	writer io.Writer
}

func (w *wrapWrite) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}

func (w *wrapWrite) WriteString(s string) (n int, err error) {
	return w.writer.Write([]byte(s))
}
