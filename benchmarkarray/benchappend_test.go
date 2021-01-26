package benchmarkarray

import "testing"

func BenchmarkAppendTail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendTail()
	}
}

func BenchmarkAppendHead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendHead()
	}
}
