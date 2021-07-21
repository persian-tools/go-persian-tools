package bank

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	// ErrBankNotFound is returned by CardInfo function
	// when bank not found for input card
	ErrBankNotFound = errors.New("bank not found")
	// ErrInvalidCard is returned by CardInfo function
	// when card has invalid format
	ErrInvalidCard = errors.New("card is not valid")
)

var bankCode = map[string]string{
	"636214": "ayandeh",
	"627412": "eghtesad-novin",
	"627381": "ansar",
	"505785": "iran-zamin",
	"622106": "parsian",
	"627884": "parsian",
	"639194": "parsian",
	"502229": "pasargad",
	"502249": "pasargad",
	"639347": "pasargad",
	"585983": "tejarat",
	"627353": "tejarat",
	"502908": "tosee-taavon",
	"207177": "export-development",
	"627648": "export-development",
	"636949": "hekmat-iranian",
	"505809": "middle-east",
	"585947": "middle-east",
	"502938": "dey",
	"589463": "refah",
	"621986": "saman",
	"589210": "sepah",
	"604932": "sepah",
	"639607": "sarmayeh",
	"639346": "sina",
	"502806": "shahr",
	"504706": "shahr",
	"603769": "saderat",
	"903769": "saderat",
	"627961": "industry-and-mine",
	"504172": "resalat",
	"606373": "mehr-iran",
	"639599": "ghavamin",
	"502910": "karafarin",
	"627488": "karafarin",
	"603770": "keshavarzi",
	"639217": "keshavarzi",
	"505416": "tourism",
	"636795": "central-bank",
	"628023": "maskan",
	"610433": "mellat",
	"991975": "mellat",
	"170019": "melli",
	"603799": "melli",
	"639370": "mehr-eqtesad",
	"627760": "post-bank",
	"628157": "tosseh",
	"606256": "melal",
	"505801": "kosar",
	"507677": "noor",
}

// CardInfo find card's bank name.
func CardInfo(card string) (string, error) {
	matched, err := regexp.MatchString("^\\d{16}$", card)
	if !matched || err != nil || !validateCard(card) {
		return "", ErrInvalidCard
	}

	bank, ok := bankCode[card[0:6]]
	if !ok {
		return "", ErrBankNotFound
	}

	return bank, nil
}

func validateCard(card string) bool {
	length := len(card)
	base1, _ := strconv.Atoi(card[1:10])
	base2, _ := strconv.Atoi(card[10:16])
	if length < 16 || base1 == 0 || base2 == 0 {
		return false
	}

	sum := 0
	var even int
	var subDigit int
	for i := 0; i < 16; i++ {
		even = 1
		if i%2 == 0 {
			even = 2
		}
		base, _ := strconv.ParseInt(card[i:i+1], 10, 0)
		subDigit = int(base) * even
		if subDigit > 9 {
			sum += subDigit - 9
		} else {
			sum += subDigit
		}
	}
	return sum%10 == 0
}
