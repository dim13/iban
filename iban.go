// Package iban implements IBAN number validator
package iban

func toInt(s string) (x []int) {
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			x = append(x, int(r-'A')+10)
		case r >= 'a' && r <= 'z':
			x = append(x, int(r-'a')+10)
		case r >= '0' && r <= '9':
			x = append(x, int(r-'0'))
		}
	}
	return
}

// IsValid IBAN
func IsValid(iban string) bool {
	n := toInt(iban)
	if len(n) < 4 || len(n) > 34 {
		return false
	}
	var x int
	for _, v := range append(n[4:], n[:4]...) {
		if v >= 10 {
			x = (x*10 + v/10) % 97
		}
		x = (x*10 + v%10) % 97
	}
	return x == 1
}
