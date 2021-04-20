package bank

import "testing"

func TestShebaCode_IsSheba(t *testing.T) {
	shebaCode := ShebaCode{"IR820540102680020817909002"}
	sheba := shebaCode.IsSheba()
	if sheba.Code != "054" || sheba.Name != "Parsian Bank" || sheba.PersianName != "بانک پارسیان" || !sheba.AccountNumberAvailable {
		t.Errorf("Result is not correct : %v", sheba)
	}

	shebaCode.Code = "IR820540102680020817909003"
	sheba2 := shebaCode.IsSheba()
	if sheba2.Code != "" || sheba2.Name != "" || sheba2.PersianName != "" || sheba2.AccountNumberAvailable {
		t.Errorf("Result is not correct : %v", sheba2)
	}
}
