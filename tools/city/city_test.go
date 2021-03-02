package city

import (
	"reflect"
	"testing"
)

func TestGetPlaceByIranNationalId(t *testing.T) {
	gpbnId := GetPlaceByIranNationalId("0499370899")
	if gpbnId.City != "شهرری" || gpbnId.Province != "تهران" || !reflect.DeepEqual(gpbnId.Codes, []string{"048", "049"}) {
		t.Errorf("Result is false : %v", gpbnId)
	}
	wrongResult := GetPlaceByIranNationalId("059499370899")
	if wrongResult.City != "" || wrongResult.Province != "" || len(wrongResult.Codes) > 0 {
		t.Errorf("Result is false : %v", wrongResult)
	}
}
