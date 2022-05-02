// Package ra1 implements modular arithmetic required for randomized algoithms.
package ra1

import (
	"math/big"
)

func IsEven(y *big.Int) bool {
	bit := big.NewInt(1)
	z := big.NewInt(0).And(y, bit)
	return z.Cmp(bit) != 0
}

// ModExp calculates x^y mod N.
func ModExp(x, y, N *big.Int) *big.Int {
	zero := big.NewInt(0)
	if y.Cmp(zero) == 0 {
		return big.NewInt(1)
	}

	yDiv2 := big.NewInt(0).Div(y, big.NewInt(2))
	z := ModExp(x, yDiv2, N)
	zSq := big.NewInt(0).Mul(z, z)

	if IsEven(y) {
		return big.NewInt(0).Mod(zSq, N)
	}

	a := big.NewInt(0).Mul(x, zSq)
	return big.NewInt(0).Mod(a, N)
}
