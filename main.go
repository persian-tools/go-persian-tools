package main

import (
	"fmt"
	bills "github.com/nimahkh/persian_tools/tools/bills"
	digits "github.com/nimahkh/persian_tools/tools/digits"
)

func main() {
	currency_param := bills.Currency{false, false}
	params := bills.BillParams{1117753200140, 12070160, currency_param, "1"}
	params2 := bills.BillParams{7748317800142, 1770160, currency_param, "1"}
	result := bills.GetBillType(params)
	amount := bills.GetCurrency(params)
	barcode := bills.GetBarCode(params)
	verify := bills.VerifyBillID(params2)

	num2word := digits.DigitToWord("۱۵۶۷۸۹")

	fmt.Printf("\n %v \n", result)
	fmt.Printf("\n %v \n", amount)
	fmt.Printf("\n %v \n", num2word)
	fmt.Printf("\n %v \n", barcode)
	fmt.Printf("\n %v \n", verify)
}
