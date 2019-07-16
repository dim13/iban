// Package iban implements IBAN number validator
package iban

import "math/big"

// IsValid IBAN
func IsValid(iban string) bool {
	if len(iban) < 4 || len(iban) > 34 {
		return false
	}
	sum := new(big.Int)
	ten := big.NewInt(10)
	for _, v := range iban[4:] + iban[:4] {
		switch {
		case v >= 'A' && v <= 'Z':
			n := int64(v - 'A' + 10)
			sum.Add(sum.Mul(sum, ten), big.NewInt(n/10))
			sum.Add(sum.Mul(sum, ten), big.NewInt(n%10))
		case v >= '0' && v <= '9':
			n := int64(v - '0')
			sum.Add(sum.Mul(sum, ten), big.NewInt(n))
		case v == ' ':
		default:
			return false
		}
	}
	return sum.Mod(sum, big.NewInt(97)).Int64() == 1
}
