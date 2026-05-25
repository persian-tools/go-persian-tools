package digit

import "strings"

var persianToEnglishMap = map[rune]rune{
	'۰': '0',
	'۱': '1',
	'۲': '2',
	'۳': '3',
	'۴': '4',
	'۵': '5',
	'۶': '6',
	'۷': '7',
	'۸': '8',
	'۹': '9',
}

func PersianToEnglish(input string) string {
	var result strings.Builder
	for _, char := range input {
		if engChar, found := persianToEnglishMap[char]; found {
			result.WriteRune(engChar)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
