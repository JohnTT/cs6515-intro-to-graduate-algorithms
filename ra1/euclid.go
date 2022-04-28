package ra1

// EuclidGCD returns the greatest common divisor of two positive integers, x and y.
func EuclidGCD(x, y int) int {
	if y == 0 {
		return x
	}
	return EuclidGCD(y, x%y)
}

// ExtEuclid returns d = gcd(x,y), α = x^-1 mod y, β = y^-1 mod x
func ExtEuclid(x, y int) (d, α, β int) {
	if y == 0 {
		return x, 1, 0
	}
	d, α_prime, β_prime := ExtEuclid(y, x%y)
	return d, β_prime, α_prime - (x/y)*β_prime
}
