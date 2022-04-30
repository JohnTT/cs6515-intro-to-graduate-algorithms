// Package ra2 implements the RSA public-key cryptosystem.
package ra2

import "math/big"

type RSAPublicKey struct {
	N big.Int
	e int
}

type RSAPrivateKey struct {
	d int
}
