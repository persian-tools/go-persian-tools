<div align="center">
	<p align="center">
		<img src="https://raw.githubusercontent.com/persian-tools/persian-tools/master/images/logo.png" width="200" />
	</p>
	<h1 align="center">Persian tools</h1>
	<p align="center">An anthology of a variety of tools for the Persian language in Golang</p>
</div>

[![CircleCI](https://circleci.com/gh/circleci/circleci-docs.svg?style=svg)](https://circleci.com/gh/circleci/circleci-docs)
[![codecov](https://codecov.io/gh/nimahkh/persian_tools/branch/master/graph/badge.svg?token=7D038RFAP9)](https://codecov.io/gh/nimahkh/persian_tools)

<hr />

#### Todos 
- [x] Bill calculator 
- [x] Digits 
- [ ] Validate Bank card number.
- [ ] Find Bank's name by Card number.
- [ ] Check Iranian Sheba(IBAN) validation and recognize bank information by sheba code. 
- [ ] Add and remove commas to numbers. 
- [ ] Find city and province name by national code(code-e Melli). 
- [ ] Validate Iranian national number(code-e Melli). 


#### How to use it?
first we have to create our request :

```go
bills "github.com/nimahkh/persian_tools/tools/bills"
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
```

