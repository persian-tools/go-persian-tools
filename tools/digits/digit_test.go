package digits

import "testing"

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
