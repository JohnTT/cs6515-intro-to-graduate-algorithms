// Package dp1 implements Fibonacci numbers (fib), Longest Increasing Subsequence (LIS),
// and Longest Common Subsequence (LCS).
package dp1

// Fib1 calculates the n-th Fibonnaci number using a recursive algorithm.
// This results in exponential running time.
func Fib1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fib1(n-1) + Fib1(n-2)
}

// Fib2 calculates the n-th Fibonnaci number using dynamic programming.
// This results in linear running time.
func Fib2(n int) int {
	F := make([]int, n+1)
	F[0] = 0
	F[1] = 1
	for i := 2; i <= n; i++ {
		F[i] = F[i-1] + F[i-2]
	}
	return F[n]
}
