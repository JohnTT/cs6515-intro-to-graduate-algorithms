package ra1

import "testing"

func TestEuclidGCD(t *testing.T) {
	if val := EuclidGCD(1, 2); val != 1 {
		t.Errorf("EuclidGCD(1, 2) = %d; want 1", val)
	}

	if val := EuclidGCD(2, 4); val != 2 {
		t.Errorf("EuclidGCD(2, 4) = %d; want 2", val)
	}
}
