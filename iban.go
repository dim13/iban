// Package iban implements IBAN number validator
package iban

import (
	"math/big"
	"unicode"
)

// IsValid IBAN
func IsValid(iban string) bool {
	if len(iban) < 4 || len(iban) > 34 {
		return false
	}
	sum := new(big.Int)
	ten := big.NewInt(10)
	for _, v := range iban[4:] + iban[:4] {
		switch v = unicode.ToUpper(v); {
		case unicode.IsLetter(v):
			n := int64(v - 'A' + 10)
			sum.Add(sum.Mul(sum, ten), big.NewInt(n/10))
			sum.Add(sum.Mul(sum, ten), big.NewInt(n%10))
		case unicode.IsDigit(v):
			n := int64(v - '0')
			sum.Add(sum.Mul(sum, ten), big.NewInt(n))
		case unicode.IsSpace(v):
			// ignore
		}
	}
	return sum.Mod(sum, big.NewInt(97)).Int64() == 1
}
