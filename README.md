<div align="center">
	<p align="center">
		<img src="https://raw.githubusercontent.com/persian-tools/persian-tools/master/images/logo.png" width="200" />
	</p>
	<h1 align="center">Persian tools</h1>
	<p align="center">An anthology of a variety of tools for the Persian language in Golang</p>
</div>

[![CircleCI](https://circleci.com/gh/circleci/circleci-docs.svg?style=svg)](https://circleci.com/gh/circleci/circleci-docs)
[![codecov](https://codecov.io/gh/persian-tools/go-persian-tools/branch/master/graph/badge.svg?token=7D038RFAP9)](https://codecov.io/gh/persian-tools/go-persian-tools)

<hr />

#### Todos

- [x] Bill calculator
- [x] Digits
- [x] Validate Bank card number.
- [x] Find Bank's name by Card number.
- [x] Check Iranian Sheba(IBAN) validation and recognize bank information by sheba code.
- [x] Add and remove commas to numbers.
- [x] Find city and province name by national code(code-e Melli).
- [x] Validate Iranian national number(code-e Melli).

#### How to use it?

first we have to create our request :

```go
bills "github.com/nimahkh/persian_tools/bills"
.
.
.
currencyParam := bills.Currency{false, false}
params := bills.BillParams{1117753200140, 12070160, currency_param, "1"}
```

And pass it

##### Bills

```go
result := bills.GetBillType(params) // تلفن ثابت
amount := bills.GetCurrency(params) //120000
barcode := bills.GetBarCode(params) //  111775320014000012070160
verify := bills.VerifyBillID(params) //true

```

##### Digits

```go
num2wordFa := digits.DigitToWord("۱۵۶۷۸۹") // صد پنجاه و شش هزار هفتصد هشتاد و نه
num2wordEn := digits.DigitToWord("156789") // صد پنجاه و شش هزار هفتصد هشتاد و نه
Negative := digits.DigitToWord("-156789") // منفی صد پنجاه و شش هزار هفتصد هشتاد و نه
englishDigits := PersianToEnglish("۱۲۳۴۵۶۷۸۹۰") // 1234567890
persianDigits := EnglishToPersian("1234567890") // ۱۲۳۴۵۶۷۸۹۰
```

#### Bank

###### CardInfo Method

Method has one entry card number and return bank name and error.

```go
bank,error := CardInfo("") // "", error
bank,error := CardInfo("345345") // "", error
bank,error := CardInfo("6219861034529008") // "", error
bank,error := CardInfo("6037701689095443") // keshavarzi, nil
bank,error := CardInfo("6219861034529007") // saman, nil

```

###### Check Iranian Sheba

The types of results are :

```go
type shebaResultHash struct {
	Name                   string
	Code                   string
	NickName               string
	PersianName            string
	AccountNumber          string
	AccountNumberAvailable bool
	FormattedAccountNumber string
	Process                func(str string) ShebaProcess
}
```

```go
shebaCode := bank.ShebaCode{"IR820540102680020817909002"}
sheba := shebaCode.IsSheba() // {Parsian Bank 054 parsian بانک پارسیان  true  0x4c69f0}
shebaCode.Code= "IR820540102680020817909003"
sheba := shebaCode.IsSheba() // {  false }
```

###### Add Remove Commas to number

```go
addComma := digits.AddCommas(14555478854)
removeComma := digits.RemoveCommas("4,555,522,212,12")

fmt.Printf("\n ADD COMMA : %v \n", addComma) // 14,555,478,854
fmt.Printf("\n REMOVE COMMA : %v \n", removeComma)// 455552221212
```

###### Get Place and city By NationalID

```go
getPlaceByIranNationalId := city.GetPlaceByIranNationalId("0499370899")
fmt.Printf("\n Result : %v \n", getPlaceByIranNationalId)

```

###### Validate Iranian national number(code-e Melli)

```go
verifyIranianNationalId := national_id.Validate("0067749828")
verifyIranianNationalIdFalse := national_id.Validate("0684159415")

fmt.Printf("\n Validate NationalID : %v \n", verifyIranianNationalId) // true
fmt.Printf("\n Validate NationalIDFalse : %v \n", verifyIranianNationalIdFalse) // false
```
