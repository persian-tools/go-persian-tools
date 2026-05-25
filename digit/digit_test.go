package digit

import (
	"testing"
)

func TestDigitToWord(t *testing.T) {
	num2word1 := DigitToWord("۱۵۶۷۸۹")
	num2word2 := DigitToWord("-10")
	num2word3 := DigitToWord("-156788")

	result := "صد پنجاه و شش هزار هفتصد هشتاد و نه"
	if num2word1 != result {
		t.Errorf("Wrong Convert : %v", result)
	}

	result2 := "منفی ده"
	if num2word2 != result2 {
		t.Errorf("Wrong Convert : %v", result2)
	}

	result3 := "منفی صد پنجاه و شش هزار هفتصد هشتاد و هشت"
	if num2word3 != result3 {
		t.Errorf("Wrong Convert : %v", result3)
	}
}

func TestAddCommas(t *testing.T) {
	addComma := AddCommas(14555478854)
	if addComma != "14,555,478,854" {
		t.Errorf("Wrong Convert : %v", addComma)
	}
}

func TestRemoveCommas(t *testing.T) {
	removeComma := RemoveCommas("4,555,522,212,12")
	if removeComma != 455552221212 {
		t.Errorf("Wrong Convert : %v", removeComma)
	}

}

func TestEnglishToPersian(t *testing.T) {
	// expect number 1234567890 to turn into ۱۲۳۴۵۶۷۸۹۰
	persianDigits := EnglishToPersian("1234567890")
	if persianDigits != "۱۲۳۴۵۶۷۸۹۰" {
		t.Errorf("Wrong Converting English Digits To Persian: %v", persianDigits)
	}

	// expect misc characters to stay the same
	persianDigitsWithMisc := EnglishToPersian("1!2@3#4$5%6^7&8*9(0)-=")
	if persianDigitsWithMisc != "۱!۲@۳#۴$۵%۶^۷&۸*۹(۰)-=" {
		t.Errorf("Wrong Converting Persian Digits To English: %v", persianDigitsWithMisc)
	}
}

func TestPersianToEnglish(t *testing.T) {
	// expect number 1234567890 to turn into ۱۲۳۴۵۶۷۸۹۰
	englishDigits := PersianToEnglish("۱۲۳۴۵۶۷۸۹۰")
	if englishDigits != "1234567890" {
		t.Errorf("Wrong Converting Persian Digits To English: %v", englishDigits)
	}
	// expect misc characters to stay the same
	englishDigitsWithMisc := PersianToEnglish("۱!۲@۳#۴$۵%۶^۷&۸*۹(۰)-=")
	if englishDigitsWithMisc != "1!2@3#4$5%6^7&8*9(0)-=" {
		t.Errorf("Wrong Converting Persian Digits To English: %v", englishDigitsWithMisc)
	}
}
