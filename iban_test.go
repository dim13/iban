package iban

import "testing"

var testCase = map[string]string{
	"Austria":        "AT611904300234573201",
	"Belgium":        "BE62510007547061",
	"Bulgaria":       "BG80BNBG96611020345678",
	"Cyprus":         "CY17002001280000001200527600",
	"Czech Republic": "CZ6508000000192000145399",
	"Denmark":        "DK5000400440116243",
	"Estonia":        "EE382200221020145685",
	"Finland":        "FI2112345600000785",
	"France":         "FR1420041010050500013M02606",
	"Germany":        "DE89370400440532013000",
	"Greece":         "GR1601101250000000012300695",
	"Hungary":        "HU42117730161111101800000000",
	"Iceland":        "IS140159260076545510730339",
	"Ireland":        "IE29AIBK93115212345678",
	"Italy":          "IT40S0542811101000000123456",
	"Latvia":         "LV80BANK0000435195001",
	"Liechtenstein":  "LI21088100002324013AA",
	"Lithuania":      "LT121000011101001000",
	"Luxembourg":     "LU280019400644750000",
	"Malta":          "MT84MALT011000012345MTLCAST001S",
	"Netherlands":    "NL39RABO0300065264",
	"Norway":         "NO9386011117947",
	"Poland":         "PL27114020040000300201355387",
	"Portugal":       "PT50000201231234567890154",
	"Romania":        "RO49AAAA1B31007593840000",
	"Slovakia":       "SK3112000000198742637541",
	"Slovenia":       "SI56191000000123438",
	"Spain":          "ES0700120345030000067890",
	"Sweden":         "SE3550000000054910000003",
	"Switzerland":    "CH9300762011623852957",
	"United Kingdom": "GB54BARC20992012345678",
}

func TestIsValid(t *testing.T) {
	for country, iban := range testCase {
		t.Run(country, func(t *testing.T) {
			if !IsValid(iban) {
				t.Errorf("%v, expected valid", iban)
			}
		})
	}
}
