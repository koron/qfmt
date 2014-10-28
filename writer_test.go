package qfmt

import (
	"io/ioutil"
	"os"
	"testing"
)

type writer struct{}

func (w *writer) Write(p []byte) (int, error) {
	return 0, nil
}

func TestWrapWriter(t *testing.T) {
	w := to_w(&writer{})
	if _, ok := w.(*wrapWriter); !ok {
		t.Errorf("writer must be wrapped")
	}
}

func TestNoWrapWriter(t *testing.T) {
	w := to_w(ioutil.Discard)
	if _, ok := w.(*wrapWriter); ok {
		t.Errorf("ioutil.Discard must not be wrapped")
	}
}

func TestNoWrapWriter2(t *testing.T) {
	w := to_w(os.Stdout)
	if _, ok := w.(*wrapWriter); ok {
		t.Errorf("os.Stdout must be wrapped")
	}
}
