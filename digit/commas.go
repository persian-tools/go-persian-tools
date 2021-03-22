package digit

import (
	"github.com/dustin/go-humanize"
	"strconv"
	"strings"
)

func AddCommas(digit int64) string {
	return humanize.Comma(digit)
}

func RemoveCommas(string string) int {
	removeStr, _ := strconv.Atoi(strings.ReplaceAll(string, ",", ""))
	return removeStr
}
