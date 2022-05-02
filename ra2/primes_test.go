package ra2

import (
	"math/big"
	"testing"
)

const n = 64

func TestIsPrime(t *testing.T) {
	var x *big.Int
	var want, val bool

	// IsPrime(2) = true.
	x = big.NewInt(2)
	val = IsPrime(x, n)
	want = true
	if val != want {
		t.Errorf("IsPrime(%d) = %v, want %v\n", x, val, want)
	}

	// IsPrime(3) = true.
	x = big.NewInt(3)
	val = IsPrime(x, n)
	want = true
	if val != want {
		t.Errorf("IsPrime(%d) = %v, want %v\n", x, val, want)
	}

	// IsPrime(4) = false.
	x = big.NewInt(4)
	val = IsPrime(x, n)
	want = false
	if val != want {
		t.Errorf("IsPrime(%d) = %v, want %v\n", x, val, want)
	}

	// IsPrime(5) = true.
	x = big.NewInt(5)
	val = IsPrime(x, n)
	want = true
	if val != want {
		t.Errorf("IsPrime(%d) = %v, want %v\n", x, val, want)
	}

	// IsPrime(7) = true.
	x = big.NewInt(5)
	val = IsPrime(x, n)
	want = true
	if val != want {
		t.Errorf("IsPrime(%d) = %v, want %v\n", x, val, want)
	}

	// IsPrime(9) = false.
	x = big.NewInt(5)
	val = IsPrime(x, n)
	want = true
	if val != want {
		t.Errorf("IsPrime(%d) = %v, want %v\n", x, val, want)
	}

	// x = 1729 is a Carmichael number.
	// IsPrime(1729) = false.
	x = big.NewInt(1729)
	val = IsPrime(x, n)
	want = false
	if val != want {
		t.Errorf("IsPrime(%d) = %v, want %v\n", x, val, want)
	}
}

func TestRandomPrime(t *testing.T) {
	p, _ := RandomPrime(n)
	t.Logf("\np = %X\n", p)
}

func TestRandomInteger(t *testing.T) {
	for i := 0; i < 1000; i++ {
		x, _ := RandomInteger(n)
		if x.Cmp(big.NewInt(1)) == -1 {
			t.Errorf("RandomInteger(%d) = %d\n", n, x)
		}
	}
}

func BenchmarkRandomInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomInteger(n)
	}
}

func BenchmarkRandomPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomPrime(n)
	}
}
