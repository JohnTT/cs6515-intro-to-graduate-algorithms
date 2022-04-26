package ra1

import "testing"

func TestModExp(t *testing.T) {
	if val := ModExp(7, 5, 23); val != 17 {
		t.Errorf("ModExp(7, 5, 23) = %d; want 17", val)
	}

	if val := ModExp(7, 25, 23); val != 21 {
		t.Errorf("ModExp(7, 25, 23) = %d; want 21", val)
	}
}
