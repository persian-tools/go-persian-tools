package bank

import (
	"testing"
)

func TestCard_CardInfo(t *testing.T) {
	card := Card{Number: "6219861034529008"}
	falseResult := card.CardInfo()
	card = Card{Number: "6037701689095443"}
	bankKeshavarzi := card.CardInfo()
	card = Card{Number: "6219861034529007"}
	bankSaman := card.CardInfo()

	if falseResult.isValid || falseResult.bankName != "" {
		t.Errorf("Result has to be false : %v", falseResult)
	}

	if !bankKeshavarzi.isValid || bankKeshavarzi.bankName != "بانک کشاورزی" {
		t.Errorf("Result is not correct : %v", bankKeshavarzi)
	}

	if !bankSaman.isValid || bankSaman.bankName != "بانک سامان" {
		t.Errorf("Result is not correct : %v", bankSaman)
	}
}

func TestShebaCode_IsSheba(t *testing.T) {
	shebaCode := ShebaCode{"IR820540102680020817909002"}
	sheba := shebaCode.IsSheba()
	if sheba.Code != "054" || sheba.Name != "Parsian Bank" || sheba.PersianName != "بانک پارسیان" || !sheba.AccountNumberAvailable {
		t.Errorf("Result is not correct : %v", sheba)
	}

	shebaCode.Code = "IR820540102680020817909003"
	sheba2 := shebaCode.IsSheba()
	if sheba2.Code != "" || sheba2.Name != "" || sheba2.PersianName != "" || sheba2.AccountNumberAvailable {
		t.Errorf("Result is not correct : %v", sheba2)
	}
}
