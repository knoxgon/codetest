package ibanpkg

/*
	Author: Volkan Güven
	github: https://github.com/knoxgon/

	Resources:
	https://usersite.datalab.eu/printclass.aspx?type=wiki&id=91772
	https://bank.codes/iban/structure
	https://en.wikipedia.org/wiki/International_Bank_Account_Number

	Wikipedia Quote:

	Validating the IBAN
	An IBAN is validated by converting it into an integer and performing a basic mod-97 operation
	(as described in ISO 7064) on it. If the IBAN is valid, the remainder equals 1.
	The algorithm of IBAN validation is as follows:

	  Check that the total IBAN length is correct as per the country. If not, the IBAN is invalid
	 Move the four initial characters to the end of the string
	 Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35
	 Interpret the string as a decimal integer and compute the remainder of that number on division by 97
	 If the remainder is 1, the check digit test is passed and the IBAN might be valid.

	Example (fictitious United Kingdom bank, sort code 12-34-56, account number 98765432):

	• IBAN:		                GB82 WEST 1234 5698 7654 32
	• Rearrange:		        W E S T12345698765432 G B82
	• Convert to integer:		3214282912345698765432161182
	• Compute remainder:		3214282912345698765432161182	mod 97 = 1

*/

import (
	"math/big"
	"strconv"
	"strings"
)

//getIbanCodes() returns a IBAN map that has the country code and corresponding length
//I.e., Sweden has iso2 SE and iban number length 22
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
		"GP": 27,
	}
}

//controlIso2AndLength controls the first 2 iban characters
//to match it with the existing iban conditions from getIbanCodes()
func controlIso2AndLength(iban string) bool {
	//Trim spaces
	iban = strings.Replace(iban, " ", "", -1)

	//If iso2 exists in the library and total length is matched and minimum length is 5
	for k, v := range getIbanCodes() {
		if k == iban[0:2] && len(iban) == v && len(iban) >= 5 {
			return true
		}
	}
	return false
}

//charToIso7064 Converts each alphanumeric characters from A to Z to base 36.
//Numbers 0 to 9 are converted to ASCII so i.e., '5' is 5
func charToIso7064(char rune) int64 {
	digit := int64(char)
	if char >= 'A' && char <= 'Z' {
		if digit > 64 && digit < 91 {
			digit -= 55
			return digit
		}
	} else if char >= '0' && char <= '9' {
		digit -= '0'
		return digit
	}
	return -1
}

/*ControlIban controls the iban number
The return type is a condition whether the given iban is valid or not.
*/
func ControlIban(iban string) bool {
	//Terminate before going any further if iban has fewer characters
	if len(iban) < 5 {
		return false
	}

	iban = strings.ToUpper(strings.Replace(iban, " ", "", -1))
	if controlIso2AndLength(iban) != true {
		return false
	}

	//Slice the countrycode and checksum
	countryCode := iban[0:4]
	//Get the rest of the bban
	bban := iban[4:]

	//Move the iso2 and checksum to the end of the string
	rotatediban := bban + countryCode

	//New string with the standard iso7064 is being written to rearrange
	rearrange := ""

	//The for loop converts every character to base36,
	//any character not meeting the requirement breaks the function
	//(Probable faulty IBAN)
	for _, c := range rotatediban {
		elem := charToIso7064(c)
		if elem == -1 {
			return false
		}
		rearrange += strconv.Itoa(int(elem))
	}

	//After converting the iban to base 36, write it back as a string
	//Since the number is too large, uint64 cannot hold it, thus SetString
	//is used to hold larger numbers
	ibanAsInt, isOk := new(big.Int).SetString(rearrange, 10)

	if !isOk {
		panic("Error happened while processing 'rearrange'")
	}

	//Since big.Int.Mod(x,y) function returns *Int, we have to have modulus 97 as a big int
	mod := big.NewInt(97)

	//The result is either 1 or non-1, anything else beside 1 is false
	result := new(big.Int).Mod(ibanAsInt, mod)

	//result is a pointer to an Int, so, converting to Int64() allows us to make comparision
	return result.Int64() == 1
}
