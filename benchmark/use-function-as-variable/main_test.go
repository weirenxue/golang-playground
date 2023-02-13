package main

import "testing"

func BenchmarkStringCombine1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringCombine1("Hello", "World")
	}
}

func BenchmarkStringCombine2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringCombine2("Hello", "World")
	}
}
