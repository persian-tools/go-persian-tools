package bills

import (
	"testing"
)

var currencyParam = Currency{false, false}
var params = BillParams{1117753200140, 12070160, currencyParam, "1"}

func TestGetBillType(t *testing.T) {
	var params2 = BillParams{9100074409151, 12908190, currencyParam, "1"}
	var params3 = BillParams{7748317800105, 1770160, currencyParam, "1"}
	var billResultType = "تلفن ثابت"
	var billResultType2 = "تلفن همراه"
	var billResultType3 = "unknown"

	bt := GetBillType(params)
	if bt != billResultType {
		t.Errorf("Type is incorrect : %v", bt)
	}

	bt2 := GetBillType(params2)
	if bt2 != billResultType2 {
		t.Errorf("Type is incorrect : %v", bt2)
	}

	bt3 := GetBillType(params3)
	if bt3 != billResultType3 {
		t.Errorf("Type is incorrect : %v", bt3)
	}
}

func TestGetCurrency(t *testing.T) {
	currencyParam := Currency{false, false}
	currencyParam2 := Currency{true, false}
	var params1 = BillParams{1117753200140, 12070160, currencyParam, "1"}
	var params2 = BillParams{1117753200140, 12070160, currencyParam2, "1"}
	var amount1 = 120000
	var amount2 = 12000

	a1 := GetCurrency(params1)
	if a1 != amount1 {
		t.Errorf("Amount is not equal : %v", a1)
	}

	a2 := GetCurrency(params2)
	if a2 != amount2 {
		t.Errorf("Amount is not equal : %v", a2)
	}

}

func TestVerifyBillID(t *testing.T) {
	var params1 = BillParams{7748317800142, 1770160, currencyParam, "1"}
	var params2 = BillParams{2234322344613, 1070189, currencyParam, "1"}

	a1 := VerifyBillID(params1)
	if a1 != true {
		t.Errorf("Wrong result : %v", a1)
	}

	a2 := VerifyBillID(params2)
	if a2 != false {
		t.Errorf("Wrong result : %v", a2)
	}
}

func TestGetBarCode(t *testing.T) {
	var params1 = BillParams{7748317800142, 1770160, currencyParam, "1"}
	var params2 = BillParams{9174639504124, 12908197, currencyParam, "1"}
	var barcode1 = "77483178001420001770160"
	var barcode2 = "917463950412400012908197"

	b1 := GetBarCode(params1)
	if b1 != barcode1 {
		t.Errorf("Barcode is not correct : %v", b1)
	}

	b2 := GetBarCode(params2)
	if b2 != barcode2 {
		t.Errorf("Barcode is not correct : %v", b2)
	}
}
