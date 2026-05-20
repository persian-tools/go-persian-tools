package persian

import (
	"regexp"
)

var faAlphabet string = `\x{0621}-\x{0628}\x{062A}-\x{063A}\x{0641}-\x{0642}\x{0644}-\x{0648}\x{064E}-\x{0651}\x{0655}\x{067E}\x{0686}\x{0698}\x{06A9}\x{06AF}\x{06BE}\x{06CC}`
var faNumber string = `\x{06F0}-\x{06F9}`

// var faShortVowels string = "َُِ"
var faMixedWithArabic string = `\x{0629}\x{0643}\x{0649}-\x{064B}\x{064D}\x{06D5}`
var faText string = faAlphabet + faNumber //+ faShortVowels
var faComplexText string = faText + faMixedWithArabic
var trimPattern string = `[\"'-+()؟\s.]`

func IsFarsi(
	text string,
	isComplex bool,
) bool {
	return IsPersian(text, isComplex)
}

func IsPersian(
	text string,
	isComplex bool,
) bool {
	regexTrimPattern := regexp.MustCompile(trimPattern)
	text = string(regexTrimPattern.ReplaceAll([]byte(text), []byte("")))

	var regex string
	if isComplex {
		regex = faComplexText
	} else {
		regex = faText
	}

	persianCompiledRegex := regexp.MustCompile("^[" + regex + "]+&")
	for _, item := range persianCompiledRegex.Split(text, -1) {
		if item != "" && item != " " {
			return false
		}

	}
	return true
}
