package digit

import (
	"github.com/mavihq/persian"
	ntw "moul.io/number-to-words"
	"strconv"
)

func DigitToWord(word string) string {
	//converter, err := word2number.NewConverter("en")
	//if err != nil {
	//	panic(err)
	//}
	//var f float64
	//f = converter.Words2Number("forty two")
	convertToEnglish, _ := strconv.Atoi(persian.ToEnglishDigits(word))
	converted := ntw.IntegerToIrIr(convertToEnglish)
	return converted
}
