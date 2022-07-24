package national_id

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

func TestValidate(t *testing.T) {
	verifyIranianNationalId := Validate("0067749828")
	verifyIranianNationalIdFalse := Validate("0684159415")
	verifyIranianNationalExceptionIdTrue := Validate("1111111111")
	verifyIranianNationalExceptionIdFalse := Validate("9999999999")

	if !verifyIranianNationalId {
		t.Errorf("Result is false : %v", verifyIranianNationalId)
	}
	if verifyIranianNationalIdFalse {
		t.Errorf("Result is false : %v", verifyIranianNationalIdFalse)
	}
	if !verifyIranianNationalExceptionIdTrue {
		t.Errorf("Result is True : %v", verifyIranianNationalExceptionIdTrue)
	}
	if verifyIranianNationalExceptionIdFalse {
		t.Errorf("Result is false : %v", verifyIranianNationalExceptionIdFalse)
	}
}
