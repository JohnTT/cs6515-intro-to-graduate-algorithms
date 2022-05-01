package ra1

import "math/big"

// EuclidGCD returns the greatest common divisor of two positive integers, x and y.
func EuclidGCD(x, y *big.Int) *big.Int {
	if y.Cmp(big.NewInt(0)) == 0 {
		return x
	}
	z := big.NewInt(0).Mod(x, y)
	return EuclidGCD(y, z)
}

// ExtEuclid returns d = gcd(x,y), α = x^-1 mod y, β = y^-1 mod x
func ExtEuclid(x, y *big.Int) (d, α, β *big.Int) {
	if y.Cmp(big.NewInt(0)) == 0 {
		return x, big.NewInt(1), big.NewInt(0)
	}

	q := big.NewInt(0).Mod(x, y)
	d, α_prime, β_prime := ExtEuclid(y, q)

	r := big.NewInt(0).Div(x, y)
	s := big.NewInt(0).Mul(r, β_prime)
	t := big.NewInt(0).Add(α_prime, s)
	return d, β_prime, t
}
