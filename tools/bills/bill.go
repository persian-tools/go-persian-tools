package bills

import (
	"strconv"
)

type Currency struct {
	Toman bool
	Rial  bool
}

type billType struct {
	bNum  int
	bType string
}

type BillBarcodeModel struct {
	billId    int
	paymentId int
}

type BillResult struct {
	amount             int
	btype              string
	barcode            string
	isValid            bool
	isValidBillId      bool
	isValidBillPayment bool
}

type BillParams struct {
	BillId    int
	PaymentId int
	Currency  Currency
	Barcode   string
}

func bills(id int) billType {
	allBills := make(map[int]billType)
	allBills[0] = billType{1, "آب"}
	allBills[1] = billType{2, "برق"}
	allBills[2] = billType{3, "گاز"}
	allBills[3] = billType{4, "تلفن ثابت"}
	allBills[4] = billType{5, "تلفن همراه"}
	allBills[5] = billType{6, "عوارض شهرداری"}
	allBills[6] = billType{8, "سازمان مالیات"}
	allBills[7] = billType{9, "جرایم راهنمایی و رانندگی"}
	allBills[10] = billType{0, "unknown"}

	return allBills[id]
}

func GetBillType(billParams BillParams) string {
	billIdString := strconv.Itoa(billParams.BillId)
	billIdLen := len(billIdString)
	billId, _ := strconv.Atoi(billIdString[billIdLen-2 : billIdLen-1])
	if billId == 0 {
		return bills(10).bType
	}
	return bills(billId - 1).bType
}

func VerifyBillID(billParams BillParams) bool {
	newBillId := strconv.Itoa(billParams.BillId)
	result := false

	if len(newBillId) < 6 {
		return false
	}

	controlBit := newBillId[len(newBillId)-1:]
	newBillId = newBillId[:len(newBillId)-1]

	c := CalTheBit(newBillId)
	controlInt, _ := strconv.Atoi(controlBit)
	result = c == controlInt

	billType := GetBillType(billParams)

	return result && billType != "unknown"
}

func GetBarCode(billParams BillParams) string {
	billID := strconv.Itoa(billParams.BillId)
	paymentID := strconv.Itoa(billParams.PaymentId)
	return billID + "000" + paymentID
}

func GetCurrency(billParams BillParams) int {
	currency := 100
	if billParams.Currency.Rial || !billParams.Currency.Toman {
		currency = 1000
	}

	payment := strconv.Itoa(billParams.PaymentId)
	payAmount, _ := strconv.Atoi(payment[0 : len(payment)-5])

	var amount = payAmount * currency

	return amount
}

func CalTheBit(num string) int {
	sum := 0
	Base := 2
	for i := 0; i < len(num); i++ {
		if Base == 8 {
			Base = 2
		}

		subString := num[len(num)-1-i : len(num)-i]
		subStringInt, _ := strconv.Atoi(subString)
		sum += subStringInt * Base
		Base++
	}
	sum %= 11
	if sum < 2 {
		sum = 0
	} else {
		sum = 11 - sum
	}
	return sum
}
