package _4_optimization

import (
	"testing"
)


func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
	}
}

var result int

func BenchmarkAddFixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = add(1, 2)
	}
}

func add(i int, i2 int) int {
	return i + i2
}
