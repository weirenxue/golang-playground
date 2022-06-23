package main

import "testing"

func BenchmarkSetSlice01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetSlice01(10000)
	}
}
func BenchmarkSetSlice02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetSlice02(10000)
	}
}
func BenchmarkSetSlice03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetSlice03(10000)
	}
}
