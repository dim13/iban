package iban

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		Country, IBAN string
	}{
		{"Austria", "AT611904300234573201"},
		{"Belgium", "BE62510007547061"},
		{"Bulgaria", "BG80BNBG96611020345678"},
		{"Cyprus", "CY17002001280000001200527600"},
		{"Denmark", "DK5000400440116243"},
		{"Estonia", "EE382200221020145685"},
		{"Finland", "FI2112345600000785"},
		{"France", "FR1420041010050500013M02606"},
		{"Germany", "DE27100777770209299700"},
		{"Great Britain", "GB29NWBK60161331926819"},
		{"Greece", "GR1601101250000000012300695"},
		{"Hungary", "HU42117730161111101800000000"},
		{"Iceland", "IS140159260076545510730339"},
		{"Ireland", "IE29AIBK93115212345678"},
		{"Italy", "IT40S0542811101000000123456"},
		{"Latvia", "LV80BANK0000435195001"},
		{"Liechtenstein", "LI21088100002324013AA"},
		{"Lithuania", "LT121000011101001000"},
		{"Luxembourg", "LU280019400644750000"},
		{"Malta", "MT84MALT011000012345MTLCAST001S"},
		{"Norway", "NO9386011117947"},
		{"Poland", "PL27114020040000300201355387"},
		{"Portugal", "PT50000201231234567890154"},
		{"Romania", "RO49AAAA1B31007593840000"},
		{"Slovak Republic", "SK3112000000198742637541"},
		{"Slovenia", "SI56191000000123438"},
		{"Spain", "ES9121000418450200051332"},
		{"Sweden", "SE3550000000054910000003"},
		{"The Czech Republic", "CZ6508000000192000145399"},
		{"The Netherlands", "NL91ABNA0417164300"},
	}
	for _, tc := range testCases {
		t.Run(tc.Country, func(t *testing.T) {
			if !IsValid(tc.IBAN) {
				t.Error(tc.IBAN, "expected valid")
			}
		})
	}
}

func TestAltFormatting(t *testing.T) {
	testCases := []string{
		"DE27100777770209299700",
		"DE27 1007 7777 0209 2997 00",
		"de27100777770209299700",
		"de27 1007 7777 0209 2997 00",
	}
	for _, tc := range testCases {
		if !IsValid(tc) {
			t.Error(tc, "want valid")
		}
	}
}

func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValid("DE27100777770209299700")
	}
}
