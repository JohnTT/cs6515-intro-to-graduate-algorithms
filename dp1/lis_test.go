package dp1

import (
	"testing"
)

func TestLIS(t *testing.T) {
	nums := []int{5, 7, 4, -3, 9, 1, 10, 4, 5, 8, 9, 3}
	if val := LIS(nums); val != 6 {
		t.Errorf("LIS(nums) = %d, want 6", val)
	}
}
