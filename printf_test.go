package qfmt

import (
	"testing"
)

func formatCheck(t *testing.T, expected, fmt string, a ...interface{}) {
	s := Sprintf(fmt, a...)
	if s != expected {
		t.Errorf("%#v isn't expaned to %#v: %#v", fmt, expected, s)
	}
}

func TestFormatConst(t *testing.T) {
	formatCheck(t, "Hello World", "Hello World")
}

func TestFormatString(t *testing.T) {
	formatCheck(t, "contents", "%s", "contents")
	formatCheck(t, "PREFIX contents SUFFIX", "PREFIX %s SUFFIX", "contents")
	formatCheck(t, "contents SUFFIX", "%s SUFFIX", "contents")
	formatCheck(t, "PREFIX contents", "PREFIX %s", "contents")
}

func TestFormatInt32(t *testing.T) {
	formatCheck(t, "1234567890", "%d", 1234567890)
	formatCheck(t, "PREFIX 1234567890 SUFFIX", "PREFIX %d SUFFIX", 1234567890)
	formatCheck(t, "1234567890 SUFFIX", "%d SUFFIX", 1234567890)
	formatCheck(t, "PREFIX 1234567890", "PREFIX %d", 1234567890)
}
