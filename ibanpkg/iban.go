package ibanpkg

/*
	Resources:
	https://usersite.datalab.eu/printclass.aspx?type=wiki&id=91772
	https://bank.codes/iban/structure
*/

import (
	"math/big"
	"strconv"
	"strings"
)

func getIbanCodes() map[string]int {
	return map[string]int{
		"AL": 28,
		"AD": 24,
		"AT": 20,
		"AZ": 28,
		"BH": 22,
		"BY": 28,
		"BE": 16,
		"BA": 20,
		"BR": 29,
		"BG": 22,
		"CR": 22,
		"HR": 21,
		"CY": 28,
		"CZ": 24,
		"DK": 18,
		"DO": 28,
		"SV": 28,
		"EE": 20,
		"FO": 18,
		"FI": 18,
		"FR": 27,
		"GE": 22,
		"DE": 22,
		"GI": 23,
		"GR": 27,
		"GL": 18,
		"GT": 28,
		"VA": 22,
		"HU": 28,
		"IS": 26,
		"IQ": 23,
		"IE": 22,
		"IL": 23,
		"IT": 27,
		"JO": 30,
		"KZ": 20,
		"XK": 20,
		"KW": 30,
		"LV": 21,
		"LB": 28,
		"LI": 21,
		"LT": 20,
		"LU": 20,
		"MK": 19,
		"MT": 31,
		"MR": 27,
		"MU": 30,
		"MD": 24,
		"MC": 27,
		"ME": 22,
		"NL": 18,
		"NO": 15,
		"PK": 24,
		"PS": 29,
		"PL": 28,
		"PT": 25,
		"QA": 29,
		"RO": 24,
		"LC": 32,
		"SM": 27,
		"ST": 25,
		"SA": 24,
		"RS": 22,
		"SC": 31,
		"SK": 24,
		"SI": 19,
		"ES": 24,
		"SE": 24,
		"CH": 21,
		"TL": 23,
		"TN": 24,
		"TR": 26,
		"UA": 29,
		"AE": 23,
		"GB": 22,
		"VG": 24,
	}
}

func controlIso2AndLength(iban string) bool {
	//Trim spaces
	iban = strings.Replace(iban, " ", "", -1)

	//If iso2 exists in the library and total length is matched
	for k, v := range getIbanCodes() {
		if k == iban[0:2] && len(iban) == v {
			return true
		}
	}
	return false
}

func charToIso7064(char rune) int64 {
	if char >= 'A' && char <= 'Z' {
		digita := int64(char)

		if digita > 64 && digita < 91 {
			digita -= 55
			return digita
		}
	} else if char >= '0' && char <= '9' {
		digit := int64(char)
		digit -= '0'
		return digit
	}
	return -1
}

/*ControlIban function controls the iban number
 */
func ControlIban(iban string) bool {
	iban = strings.ToUpper(strings.Replace(iban, " ", "", -1))

	if controlIso2AndLength(iban) != true {
		return false
	}

	countryCode := iban[0:4]
	bban := iban[4:]

	rotatediban := bban + countryCode

	rearrange := ""

	for _, c := range rotatediban {
		elem := charToIso7064(c)
		if elem != -1 {
			rearrange += strconv.Itoa(int(elem))
		}
	}

	ibanAsInt, isOk := new(big.Int).SetString(rearrange, 10)

	if !isOk {
		panic("Error happened")
	}

	mod := big.NewInt(97)

	result := new(big.Int).Mod(ibanAsInt, mod)

	return result.Int64() == 1
}
