package bank

import (
	"strconv"
)

type response struct {
	bankName string
	isValid  bool
}

type Card struct {
	Number string
}

func (c *Card) CardInfo() response {
	verify := c.verifyCardNum()
	bankName := c.getBankName()
	return response{
		bankName: bankName,
		isValid:  verify,
	}
}

func (c *Card) getBankName() string {
	digits := c.Number
	length := len(digits)
	if !c.verifyCardNum() {
		return ""
	}
	if length == 16 {
		findBank := bankCode{}
		code := digits[0:6]
		for _, v := range getBankCodes() {
			if v.code == code {
				findBank = v
			}
		}

		if findBank.name != "" {
			return findBank.name
		}
	}
	return ""
}

func (c *Card) verifyCardNum() bool {
	digitsResult := c.Number

	length := len(digitsResult)
	base1, _ := strconv.Atoi(digitsResult[1:10])
	base2, _ := strconv.Atoi(digitsResult[10:16])
	if length < 16 || base1 == 0 || base2 == 0 {
		return false
	}

	sum := 0
	var even int
	var subDigit int
	for i := 0; i < 16; i++ {
		even = 1
		if i%2 == 0 {
			even = 2
		}
		base, _ := strconv.ParseInt(digitsResult[i:i+1], 10, 0)
		subDigit = int(base) * even
		if subDigit > 9 {
			sum += subDigit - 9
		} else {
			sum += subDigit
		}
	}
	return sum%10 == 0
}
