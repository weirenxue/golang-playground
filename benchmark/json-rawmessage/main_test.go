package main

import "testing"

func BenchmarkUnmarshal01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unmarshal01(100)
	}
}

func BenchmarkUnmarshal02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unmarshal02(100)
	}
}

func BenchmarkUnmarshal03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unmarshal03(100)
	}
}

func BenchmarkUnmarshal04(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unmarshal04(100)
	}
}

func BenchmarkUnmarshal05(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unmarshal05(100)
	}
}

func BenchmarkUnmarshal06(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unmarshal06(100)
	}
}
