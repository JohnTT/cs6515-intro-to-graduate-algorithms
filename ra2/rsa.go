// Package ra2 implements the RSA public-key cryptosystem.
package ra2

import "math/big"

type RSAPublicKey struct {
	N *big.Int
	e *big.Int
}

type RSAPrivateKey struct {
	d *big.Int
}

type RSAKey struct {
	RSAPublicKey
	RSAPrivateKey
	n int
}

// NewRSAKey returns an RSA public/private key pair with a bit
// strength of n.
func NewRSAKey(n int) *RSAKey {
	// TODO: generate N, e, and d.
	N := big.NewInt(0)
	e := big.NewInt(0)
	public := RSAPublicKey{N, e}

	d := big.NewInt(0)
	private := RSAPrivateKey{d}

	key := &RSAKey{public, private, n}
	return key
}
