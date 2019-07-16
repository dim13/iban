// Package iban implements IBAN number validator
package iban

// IsValid IBAN
func IsValid(iban string) bool {
	if len(iban) < 4 || len(iban) > 34 {
		return false
	}
	var x rune
	for _, v := range iban[4:] + iban[:4] {
		switch {
		case v >= 'A' && v <= 'Z':
			n := v - 'A' + 10
			x = (x*10 + n/10) % 97
			x = (x*10 + n%10) % 97
		case v >= '0' && v <= '9':
			n := v - '0'
			x = (x*10 + n) % 97
		}
	}
	return x == 1
}
