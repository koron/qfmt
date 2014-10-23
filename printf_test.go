package qfmt

import (
	"fmt"
	"io/ioutil"
	"testing"
)

////////////////////////////////////////////////////////////////////////////
// Benchmark

func BenchmarkQfmtFprintfConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fprintf(ioutil.Discard, "Hello World")
	}
}

func BenchmarkFmtFprintfConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "Hello World")
	}
}

func BenchmarkQfmtFprintfString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fprintf(ioutil.Discard, "PREFIX %s SUFFIX", "contents")
	}
}

func BenchmarkFmtFprintfString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "PREFIX %s SUFFIX", "contents")
	}
}

func BenchmarkQfmtFprintfInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fprintf(ioutil.Discard, "PREFIX %d SUFFIX", 1234567890)
	}
}

func BenchmarkFmtFprintfInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "PREFIX %d SUFFIX", 1234567890)
	}
}

func BenchmarkQfmtSprintfConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintf("Hello World")
	}
}

func BenchmarkFmtSprintfConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello World")
	}
}

func BenchmarkQfmtSprintfString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintf("PREFIX %s SUFFIX", "contents")
	}
}

func BenchmarkFmtSprintfString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("PREFIX %s SUFFIX", "contents")
	}
}

func BenchmarkQfmtSprintfInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintf("PREFIX %d SUFFIX", 1234567890)
	}
}

func BenchmarkFmtSprintfInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("PREFIX %d SUFFIX", 1234567890)
	}
}

////////////////////////////////////////////////////////////////////////////
// Tests

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

func TestFormatInt(t *testing.T) {
	formatCheck(t, "1234567890", "%d", 1234567890)
	formatCheck(t, "PREFIX 1234567890 SUFFIX", "PREFIX %d SUFFIX", 1234567890)
	formatCheck(t, "1234567890 SUFFIX", "%d SUFFIX", 1234567890)
	formatCheck(t, "PREFIX 1234567890", "PREFIX %d", 1234567890)
}
