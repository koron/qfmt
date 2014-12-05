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

func BenchmarkQfmtFprintfFourString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fprintf(ioutil.Discard, "%s.%s.%s.%s", "192", "168", "10", "254")
	}
}

func BenchmarkFmtFprintfFourString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fprintf(ioutil.Discard, "%s.%s.%s.%s", "192", "168", "10", "254")
	}
}

func BenchmarkQfmtFprintfTenString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fprintf(ioutil.Discard, "%s.%s.%s.%s.%s.%s.%s.%s.%s.%s", "000",
			"111", "222", "333", "444", "555", "666", "777", "888", "999")
	}
}

func BenchmarkFmtFprintfTenString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%s.%s.%s.%s.%s.%s.%s.%s.%s.%s", "000",
			"111", "222", "333", "444", "555", "666", "777", "888", "999")
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

func BenchmarkQfmtFprintfIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fprintf(ioutil.Discard, "%d.%d.%d.%d", 192, 168, 10, 254)
	}
}

func BenchmarkFmtFprintfIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%d.%d.%d.%d", 192, 168, 10, 254)
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

func BenchmarkQfmtSprintfFourString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintf("%s.%s.%s.%s", "192", "168", "10", "254")
	}
}

func BenchmarkFmtSprintfFourString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintf("%s.%s.%s.%s", "192", "168", "10", "254")
	}
}

func BenchmarkQfmtSprintfTenString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintf("%s.%s.%s.%s.%s.%s.%s.%s.%s.%s", "000",
			"111", "222", "333", "444", "555", "666", "777", "888", "999")
	}
}

func BenchmarkFmtSprintfTenString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%s.%s.%s.%s.%s.%s.%s.%s.%s.%s", "000",
			"111", "222", "333", "444", "555", "666", "777", "888", "999")
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

func BenchmarkQfmtSprintfIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintf("%d.%d.%d.%d", 192, 168, 10, 254)
	}
}

func BenchmarkFmtSprintfIPv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d.%d.%d.%d", 192, 168, 10, 254)
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

func TestFormatIPv4(t *testing.T) {
	formatCheck(t, "192.168.10.254", "%d.%d.%d.%d", 192, 168, 10, 254)
}
