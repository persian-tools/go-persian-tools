package bank

import (
	"errors"
	"math/big"
	"regexp"
	"strconv"
)

type ShebaCode struct {
	Code string
}

func (s ShebaCode) IsSheba() shebaResultHash {
	shebaCode := s.Code

	if !s.validate() {
		return shebaResultHash{}
	}

	regexTest, _ := regexp.Compile(`(IR[0-9]{2}([0-9]{3})[0-9]{19})`)
	code := regexTest.FindStringSubmatch(shebaCode)
	bank := shebaHashTable(code[2])

	if bank.Name == "" {
		return shebaResultHash{}
	}

	return bank
}

func iso7064Mod97_10(iban string) error {
	remainder := iban
	modVal := new(big.Int).SetInt64(97)
	bigVal, success := new(big.Int).SetString(remainder, 10)
	if !success {
		return errors.New("IBAN has incorrect check digits")
	}
	resVal := new(big.Int).Mod(bigVal, modVal)
	if resVal.Int64() != 1 {
		return errors.New("IBAN is not correct")
	}
	return nil
}

func (s ShebaCode) validate() bool {
	shebaCode := s.Code

	if len(shebaCode) != 26 {
		return false
	}
	regexPattern, _ := regexp.Compile(`(IR[0-9]{24})`)

	if !regexPattern.MatchString(shebaCode) {
		return false
	}

	d1 := []rune(shebaCode)[0] - 65 + 10
	d2 := []rune(shebaCode)[1] - 65 + 10

	newStr := shebaCode[4:]
	newStr += strconv.Itoa(int(d1)) + strconv.Itoa(int(d2)) + shebaCode[2:4]

	if iso7064Mod97_10(newStr) != nil {
		return false
	}

	return true
}
