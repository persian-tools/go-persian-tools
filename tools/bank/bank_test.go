package bank

import "testing"

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
