package ra1

// EuclidGCD returns the greatest common divisor of two positive integers, x and y.
func EuclidGCD(x, y int) int {
	if y == 0 {
		return x
	}
	return EuclidGCD(y, x%y)
}
