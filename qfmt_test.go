package qfmt

import "testing"

func BenchmarkQfmtNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = New("PREFIX %s SUFFIX")
	}
}
