package ra2

import (
	"errors"
	"math/big"
	"math/rand"
	"time"

	"github.com/JohnTT/cs6515-intro-to-graduate-algorithms/ra1"
)

func RandomInteger(n int) (*big.Int, error) {
	p := big.NewInt(0)
	if n%64 != 0 {
		return p, errors.New("n must be a multiple of 64")
	}

	// Generate a new n-bit integer p, 64 bits at a time.
	for i := 0; i < n/64; i++ {
		rand.Seed(time.Now().UnixNano())

		p.Lsh(p, 64)
		var y *big.Int
		// Most significant word must be positive.
		if i == 0 {
			y = big.NewInt(rand.Int63())
		} else {
			y = big.NewInt(int64(rand.Uint64()))
		}
		p.Add(p, y)
	}

	return p, nil
}

func IsPrime(r *big.Int, n int) bool {
	// Only consider integers >= 1.
	if r.Cmp(big.NewInt(1)) == -1 {
		return false
	}

	// Choose z_0, z_1, ..., z_k randomly from {1, 2, ..., r-1}
	// Let q = r-1
	const k = 100
	q := big.NewInt(0).Sub(r, big.NewInt(1))
	z := make([]*big.Int, k)

	for i := 0; i < k; i++ {
		// Randomly choose z_i < r.
		z[i], _ = RandomInteger(n)
		z[i].Mod(z[i], q)
		z[i].Add(z[i], big.NewInt(1))

		// Fermat's test:
		// For all i, if z_i ^ (r-1) ≡ 1 mod r, then r is prime.
		// Hence, if any z_i ^ (r-1) !≡ 1 mod r, then r is composite.
		m := ra1.ModExp(z[i], q, r)
		if m.Cmp(big.NewInt(1)) != 0 {
			return false
		}

		// TODO: fixme test for Carmichael numbers
		// // However, Fermat's test does not work for Carmichael numbers.
		// // If repeated squaring yields a non-trivial square root of 1,
		// // then r is composite.
		// //
		// // First, factorize q such that:
		// // q = 2^e * odd
		// e := 0
		// odd := big.NewInt(0)
		// odd.Add(odd, q)
		// for ra1.IsEven(odd) {
		// 	e++
		// 	odd.Div(odd, big.NewInt(2))
		// }

		// // Ignore odd == 1
		// if odd.Cmp(big.NewInt(1)) == 0 {
		// 	continue
		// }

		// // 1) Calculate s = z^odd mod r
		// // 2) Calculate s^2 mod r, up to e times.
		// // 3) If any result in (2) is 1, then r is composite.
		// s := ra1.ModExp(z[i], odd, r)
		// for ii := 0; ii < e; ii++ {
		// 	s = ra1.ModExp(s, big.NewInt(2), r)
		// 	if s.Cmp(big.NewInt(1)) == 0 {
		// 		return false
		// 	}
		// }
	}
	return true
}

// RandomPrime returns a random n-bit prime number.
func RandomPrime(n int) (*big.Int, error) {
	// Generate a new n-bit integer p.
	// Repeat until p is a prime number.
	var p *big.Int
	var err error
	for {
		p, err = RandomInteger(n)
		if err != nil {
			return p, err
		}
		if IsPrime(p, n) {
			break
		}
	}

	return p, nil
}
