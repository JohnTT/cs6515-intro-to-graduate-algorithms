package dp1

// LIS returns the length of the longest increasing subsequence in a slice of integers.
func LIS(a []int) int {
	n := len(a)
	L := make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = 1
		for j := 0; j < i; j++ {
			if a[j] < a[i] && L[i] < 1+L[j] {
				L[i] = 1 + L[j]
			}
		}
	}

	max := 1
	for i := 1; i < n; i++ {
		if L[i] > max {
			max = L[i]
		}
	}
	return max
}
