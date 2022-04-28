package ra1

import "testing"

func TestEuclidGCD(t *testing.T) {
	// Test gcd(1, 2) is 1
	if val := EuclidGCD(1, 2); val != 1 {
		t.Errorf("EuclidGCD(1, 2) = %d; want 1", val)
	}

	// Test gcd(2, 4) is 2
	if val := EuclidGCD(2, 4); val != 2 {
		t.Errorf("EuclidGCD(2, 4) = %d; want 2", val)
	}
}

func TestExtEuclid(t *testing.T) {
	// Test 1^-1 ≡ 1 mod 14
	if d, α, β := ExtEuclid(1, 14); d != 1 || α != 1 {
		t.Errorf("ExtEuclid(1, 2) d, α, β = %d, %d, %d; want 1, 1, _", d, α, β)
	}

	// Test gcd(2, 14) = 2
	if d, _, _ := ExtEuclid(2, 4); d != 2 {
		t.Errorf("ExtEuclid(2, 4) d = %d; want 2", d)
	}

	// Test 3^-1 ≡ 5 mod 14
	if d, α, β := ExtEuclid(3, 14); d != 1 || α != 5 {
		t.Errorf("ExtEuclid(3, 14) d, α, β = %d, %d, %d; want 1, 5, _", d, α, β)
	}

	// Test gcd(4, 14) = 2
	if d, _, _ := ExtEuclid(4, 14); d != 2 {
		t.Errorf("ExtEuclid(4, 14) d = %d; want 2", d)
	}

	// Test 5^-1 ≡ 3 mod 14
	if d, α, β := ExtEuclid(5, 14); d != 1 || α != 3 {
		t.Errorf("ExtEuclid(5, 14) d, α, β = %d, %d, %d; want 1, 3, _", d, α, β)
	}

	// Test 7^-1 ≡ 103 mod 360
	if d, α, β := ExtEuclid(7, 360); d != 1 || α != 103 {
		t.Errorf("ExtEuclid(5, 14) d, α, β = %d, %d, %d; want 1, 103, _", d, α, β)
	}
}
