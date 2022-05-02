package ra1

import (
	"math/big"
	"testing"
)

func TestIsEven(t *testing.T) {
	var x *big.Int
	var val bool
	var want bool

	// 2 is even
	x = big.NewInt(2)
	val = IsEven(x)
	want = true
	if val != want {
		t.Errorf("IsEven(%d) = %v, want %v", x, val, want)
	}

	// 22222 is even
	x = big.NewInt(22222)
	val = IsEven(x)
	want = true
	if val != want {
		t.Errorf("IsEven(%d) = %v, want %v", x, val, want)
	}

	// 3 is odd
	x = big.NewInt(3)
	val = IsEven(x)
	want = false
	if val != want {
		t.Errorf("IsEven(%d) = %v, want %v", x, val, want)
	}

	// 33333 is odd
	x = big.NewInt(33333)
	val = IsEven(x)
	want = false
	if val != want {
		t.Errorf("IsEven(%d) = %v, want %v", x, val, want)
	}
}

func TestModExp(t *testing.T) {
	var x, y, N, want *big.Int

	// ModExp(7, 5, 23) = 17
	x = big.NewInt(7)
	y = big.NewInt(5)
	N = big.NewInt(23)
	want = big.NewInt(17)
	if val := ModExp(x, y, N); val.Cmp(want) != 0 {
		t.Errorf("ModExp(%d, %d, %d) = %d; want %d", x, y, N, val, want)
	}

	// ModExp(7, 25, 23) = 21
	x = big.NewInt(7)
	y = big.NewInt(25)
	N = big.NewInt(23)
	want = big.NewInt(21)
	if val := ModExp(x, y, N); val.Cmp(want) != 0 {
		t.Errorf("ModExp(%d, %d, %d) = %d; want %d", x, y, N, val, want)
	}
}
