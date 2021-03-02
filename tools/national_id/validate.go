package national_id

import (
	"regexp"
	"strconv"
)

func Validate(code string) bool {
	regexTest, _ := regexp.Compile(`(^\d{10}$)`)
	match := regexTest.MatchString(code)

	if !match {
		return false
	}
	code = ("0000" + code)[len(code)+4-10:]
	stringCode, _ := strconv.Atoi(code[3:6])
	if stringCode == 0 {
		return false
	}
	lastNumber, _ := strconv.Atoi(code[9:10])
	sum := 0

	for i := 0; i < 9; i++ {
		intCode, _ := strconv.Atoi(code[i : i+1])
		sum += intCode * (10 - i)
	}
	sum = sum % 11
	result := false
	if sum < 2 && lastNumber == sum || sum >= 2 && lastNumber == 11-sum {
		result = true
	}
	return result
}
