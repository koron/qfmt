package qfmt

import (
	"testing"
)

func TestFormatConst(t *testing.T) {
	var s string
	s = Sprintf("Hello World")
	if "Hello World" != s {
		t.Errorf(`TestFormatConst is failed: %#v`, s)
	}
}

func TestFormatString(t *testing.T) {
	var s string
	s = Sprintf("%s", "foobar")
	if "foobar" != s {
		t.Errorf(`%%s is not "foobar": %#v`, s)
	}
}
