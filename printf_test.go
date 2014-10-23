package qfmt

import (
	"testing"
)

func TestFormatConst(t *testing.T) {
	var s string
	s = Sprintf("Hello World")
	if "Hello World" != s {
		t.Error(`TestFormatConst is failed: ` + s)
	}
}

func TestFormatString(t *testing.T) {
	var s string
	s = Sprintf("%s", "foobar")
	if "foobar" != s {
		t.Error(`%s is not "foobar": ` + s)
	}
}
