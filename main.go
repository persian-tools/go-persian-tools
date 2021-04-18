package main

import (
	"fmt"
	"github.com/nimahkh/persian_tools/bank"
	bills "github.com/nimahkh/persian_tools/bill"
	digits "github.com/nimahkh/persian_tools/digit"
	"github.com/nimahkh/persian_tools/national_id"
)

func main() {
	currency_param := bills.Currency{false, false}
	params := bills.BillParams{1117753200140, 12070160, currency_param, "1"}
	params2 := bills.BillParams{7748317800142, 1770160, currency_param, "1"}
	result := bills.GetBillType(params)
	amount := bills.GetCurrency(params)
	barcode := bills.GetBarCode(params)
	verify := bills.VerifyBillID(params2)
	verifyIranianNationalId := national_id.Validate("0067749828")
	verifyIranianNationalIdFalse := national_id.Validate("0684159415")

	num2word := digits.DigitToWord("۱۵۶۷۸۹")

	card := bank.Card{Number: "6037701689095443"}
	shebaCode := bank.ShebaCode{"IR820540102680020817909002"}
	bankInfo := card.CardInfo()
	sheba := shebaCode.IsSheba()
	getPlaceByIranNationalId := national_id.GetPlaceByIranNationalId("059499370899")
	addComma := digits.AddCommas(14555478854)
	removeComma := digits.RemoveCommas("4,555,522,212,12")
	fmt.Printf("\n %v \n", result)
	fmt.Printf("\n %v \n", amount)
	fmt.Printf("\n %v \n", num2word)
	fmt.Printf("\n %v \n", barcode)
	fmt.Printf("\n %v \n", verify)
	fmt.Printf("\n Verify : %v \n", bankInfo)
	fmt.Printf("\n SHEBA : %v \n", sheba)
	fmt.Printf("\n ADD COMMA : %v \n", addComma)
	fmt.Printf("\n REMOVE COMMA : %v \n", removeComma)
	fmt.Printf("\n Cities : %v \n", getPlaceByIranNationalId)
	fmt.Printf("\n Validate NationalID : %v \n", verifyIranianNationalId)
	fmt.Printf("\n Validate NationalIDFalse : %v \n", verifyIranianNationalIdFalse)
}
