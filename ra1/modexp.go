// Package ra1 implements modular exponentiation using repeated squaring.
package ra1

func isEven(y int) bool {
	return y&1 == 0
}

func ModExp(x, y, N int) int {
	if y == 0 {
		return 1
	}

	z := ModExp(x, y/2, N)

	if isEven(y) {
		return (z * z) % N
	}

	return (x * z * z) % N
}
