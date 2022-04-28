package dp1

import (
	"testing"
)

func TestFib1(t *testing.T) {
	if val := Fib1(5); val != 5 {
		t.Errorf("Fib1(5) = %d, want 5", val)
	}
}

func TestFib2(t *testing.T) {
	if val := Fib2(5); val != 5 {
		t.Errorf("Fib2(5) = %d, want 5", val)
	}
}

func BenchmarkFib1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib1(30)
	}
}

func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib2(30)
	}
}
