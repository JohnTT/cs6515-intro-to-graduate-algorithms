package ra1

import (
	"math/big"
	"testing"
)

func TestEuclidGCD(t *testing.T) {
	var x, y, gcd *big.Int

	// gcd(42, 56) = 14
	x = big.NewInt(42)
	y = big.NewInt(56)
	gcd = big.NewInt(14)
	if val := EuclidGCD(x, y); val.Cmp(gcd) != 0 {
		t.Errorf("EuclidGCD(%d, %d) = %d; want %d", x, y, val, gcd)
	}

	// gcd(461952, 116298) = 18
	x = big.NewInt(461952)
	y = big.NewInt(116298)
	gcd = big.NewInt(18)
	if val := EuclidGCD(x, y); val.Cmp(gcd) != 0 {
		t.Errorf("EuclidGCD(%d, %d) = %d; want %d", x, y, val, gcd)
	}

	// gcd(7966496, 314080416) = 32
	x = big.NewInt(7966496)
	y = big.NewInt(314080416)
	gcd = big.NewInt(32)
	if val := EuclidGCD(x, y); val.Cmp(gcd) != 0 {
		t.Errorf("EuclidGCD(%d, %d) = %d; want %d", x, y, val, gcd)
	}

	// gcd(24826148, 45296490) = 526
	x = big.NewInt(24826148)
	y = big.NewInt(45296490)
	gcd = big.NewInt(526)
	if val := EuclidGCD(x, y); val.Cmp(gcd) != 0 {
		t.Errorf("EuclidGCD(%d, %d) = %d; want %d", x, y, val, gcd)
	}
}

func TestExtEuclid(t *testing.T) {
	var x, y, gcd, inv *big.Int

	// gcd(1, 14) = 1
	// inv(1, 14) = 1
	x = big.NewInt(1)
	y = big.NewInt(14)
	gcd = big.NewInt(1)
	inv = big.NewInt(1)
	if d, α, β := ExtEuclid(x, y); d.Cmp(gcd) != 0 || α.Cmp(inv) != 0 {
		t.Errorf("ExtEuclid(%d, %d) d, α, β = %d, %d, %d; want %d, %d, _", x, y, d, α, β, gcd, inv)
	}

	// gcd(2, 14) = 2
	x = big.NewInt(2)
	y = big.NewInt(14)
	gcd = big.NewInt(2)
	if d, α, β := ExtEuclid(x, y); d.Cmp(big.NewInt(2)) != 0 {
		t.Errorf("ExtEuclid(%d, %d) d, α, β = %d, %d, %d; want %d, _, _", x, y, d, α, β, gcd)
	}

	// gcd(3, 14) = 1
	// inv(3, 14) = 5
	x = big.NewInt(3)
	y = big.NewInt(14)
	gcd = big.NewInt(1)
	inv = big.NewInt(5)
	if d, α, β := ExtEuclid(x, y); d.Cmp(gcd) != 0 || α.Cmp(inv) != 0 {
		t.Errorf("ExtEuclid(%d, %d) d, α, β = %d, %d, %d; want %d, %d, _", x, y, d, α, β, gcd, inv)
	}

	// gcd(4, 14) = 2
	x = big.NewInt(2)
	y = big.NewInt(14)
	gcd = big.NewInt(2)
	if d, α, β := ExtEuclid(x, y); d.Cmp(gcd) != 0 {
		t.Errorf("ExtEuclid(%d, %d) d, α, β = %d, %d, %d; want %d, _, _", x, y, d, α, β, gcd)
	}

	// gcd(5, 14) = 1
	// inv(5, 14) = 3
	x = big.NewInt(5)
	y = big.NewInt(14)
	gcd = big.NewInt(1)
	inv = big.NewInt(3)
	if d, α, β := ExtEuclid(x, y); d.Cmp(gcd) != 0 || α.Cmp(inv) != 0 {
		t.Errorf("ExtEuclid(%d, %d) d, α, β = %d, %d, %d; want %d, %d, _", x, y, d, α, β, gcd, inv)
	}

	// gcd(7, 360) = 1
	// inv(7, 360) = 103
	x = big.NewInt(7)
	y = big.NewInt(360)
	gcd = big.NewInt(1)
	inv = big.NewInt(103)
	if d, α, β := ExtEuclid(x, y); d.Cmp(gcd) != 0 || α.Cmp(inv) != 0 {
		t.Errorf("ExtEuclid(%d, %d) d, α, β = %d, %d, %d; want %d, %d, _", x, y, d, α, β, gcd, inv)
	}
}
