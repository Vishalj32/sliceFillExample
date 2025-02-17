package main

import (
	"testing"
)

func Benchmark_FillSliceIndex(b *testing.B) {
	slice := make([]byte, 73437)
	for i := 0; i < b.N; i++ {
		FillSliceIndex(slice, 65)
	}
}

func Benchmark_FillSliceRange(b *testing.B) {
	slice := make([]byte, 73437)
	for i := 0; i < b.N; i++ {
		FillSliceRange(slice, 66)
	}
}

func Benchmark_FillSliceCopyTrick(b *testing.B) {
	slice := make([]byte, 73437)
	for i := 0; i < b.N; i++ {
		FillSliceCopyTrick(slice, 67)
	}
}

func Benchmark_FillSlicePatternCopyTrick(b *testing.B) {
	slice := make([]byte, 73437)
	pattern := []byte{0xde, 0xad, 0xbe, 0xef}
	for i := 0; i < b.N; i++ {
		FillSlicePatternCopyTrick(slice, pattern)
	}
}
