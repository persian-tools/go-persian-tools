package city

import (
	"strings"
)

type IPlaceByNationalId struct {
	Codes    []string
	City     string
	Province string
}

func getAllCities(code string) []nationalCode {
	var allCities []nationalCode
	for _, s := range getNationalCodes() {
		findCity := strings.Contains(s.code, code)
		if findCity {
			allCities = append(allCities, s)
		}
	}
	return allCities
}

func getProvince(allCities []nationalCode) provinceCode {
	var findProvince provinceCode
	for _, s := range getProvincesCode() {
		if s.code == allCities[0].parentCode {
			findProvince = s
		}
	}
	return findProvince
}

func GetPlaceByIranNationalId(nationalId string) IPlaceByNationalId {
	if len(nationalId) == 10 {
		code := nationalId[0:3]
		allCities := getAllCities(code)

		if len(allCities) > 0 {
			findProvince := getProvince(allCities)
			codeString := allCities[0].code

			var makeCodes []string
			if strings.Contains(codeString, "-") {
				sliceCode := strings.Split(codeString, "-")
				for _, c := range sliceCode {
					makeCodes = append(makeCodes, c)
				}
			} else {
				makeCodes = append(makeCodes, codeString)
			}

			return IPlaceByNationalId{
				City:     allCities[0].city,
				Province: findProvince.city,
				Codes:    makeCodes,
			}
		}
		return IPlaceByNationalId{}
	}

	return IPlaceByNationalId{}
}
