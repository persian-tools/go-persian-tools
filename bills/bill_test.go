package bills

import (
	"testing"
)

var currencyParam = Currency{}
var params = BillParams{1117753200140, 12070160, currencyParam, "1"}

func TestGetBillType(t *testing.T) {

	var billResultType = "تلفن ثابت"
	bt := GetBillType(params)
	if bt != billResultType {
		t.Errorf("Type is incorrect : %v", bt)
	}

	params.PaymentId = 12908190
	params.BillId = 9100074409151
	var params2 = params
	var billResultType2 = "تلفن همراه"
	bt2 := GetBillType(params2)
	if bt2 != billResultType2 {
		t.Errorf("Type is incorrect : %v", bt2)
	}

	params.PaymentId = 1770160
	params.BillId = 7748317800105
	var params3 = params
	var billResultType3 = "unknown"
	bt3 := GetBillType(params3)
	if bt3 != billResultType3 {
		t.Errorf("Type is incorrect : %v", bt3)
	}
}

func TestGetCurrency(t *testing.T) {
	var amount1 = 17000
	params.Currency = currencyParam

	a1 := GetCurrency(params)
	if a1 != amount1 {
		t.Errorf("Amount is not equal : %v", a1)
	}

	currencyParam.Toman = true
	params.Currency = currencyParam
	var amount2 = 1700
	a2 := GetCurrency(params)
	if a2 != amount2 {
		t.Errorf("Amount is not equal : %v", a2)
	}

}

func TestVerifyBillID(t *testing.T) {
	params.BillId = 7748317800142
	params.PaymentId = 1770160

	a1 := VerifyBillID(params)
	if a1 != true {
		t.Errorf("Wrong result : %v", a1)
	}

	params.BillId = 2234322344613
	params.PaymentId = 1070189
	a2 := VerifyBillID(params)
	if a2 != false {
		t.Errorf("Wrong result : %v", a2)
	}
}

func TestGetBarCode(t *testing.T) {
	params.BillId = 7748317800142
	params.PaymentId = 1770160
	var barcode1 = "77483178001420001770160"
	b1 := GetBarCode(params)
	if b1 != barcode1 {
		t.Errorf("Barcode is not correct : %v", b1)
	}

	var barcode2 = "917463950412400012908197"
	params.BillId = 9174639504124
	params.PaymentId = 12908197
	b2 := GetBarCode(params)
	if b2 != barcode2 {
		t.Errorf("Barcode is not correct : %v", b2)
	}
}
