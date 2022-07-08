package main

import "testing"

func BenchmarkUnmarshallWithoutRawMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnmarshallWithoutRawMessage(10000)
	}
}

func BenchmarkUnmarshallWithRawMessage01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnmarshallWithRawMessage01(10000)
	}
}

func BenchmarkUnmarshallWithRawMessage02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnmarshallWithRawMessage02(10000)
	}
}
