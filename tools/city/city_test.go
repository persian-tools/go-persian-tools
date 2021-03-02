package city

import (
	"reflect"
	"testing"
)

func TestGetPlaceByIranNationalId(t *testing.T) {
	gpbnId := GetPlaceByIranNationalId("0499370899")
	if gpbnId.city != "شهرری" || gpbnId.province != "تهران" || !reflect.DeepEqual(gpbnId.codes, []string{"048", "049"}) {
		t.Errorf("Result is false : %v", gpbnId)
	}
	wrongResult := GetPlaceByIranNationalId("059499370899")
	if wrongResult.city != "" || wrongResult.province != "" || !reflect.DeepEqual(wrongResult.codes, []string{}) {
		t.Errorf("Result is false : %v", wrongResult)
	}
}
