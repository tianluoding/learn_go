package stringtest

import "testing"

func BenchmarkAddString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddString()
	}
}

func BenchmarkJoinString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JoinString()
	}
}
