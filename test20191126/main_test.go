package main

import "testing"

func BenchmarkMMMMain(b *testing.B) {
	b.N = 100
}
