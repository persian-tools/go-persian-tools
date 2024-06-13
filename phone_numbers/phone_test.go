package phonenumbers

import (
	"reflect"
	"testing"
)

func TestIsPhoneValid(t *testing.T) {
	tests := []struct {
		phoneNumber string
		isValid     bool
	}{
		{"9122221811", true},
		{"09122221811", true},
		{"+989122221811", true},
		{"12903908", false},
		{"901239812390812908", false},
	}

	for _, tt := range tests {
		ok := IsPhoneValid(tt.phoneNumber)
		if ok != tt.isValid {
			t.Errorf("IsPhoneValid(%s) got %t, expected %t", tt.phoneNumber, ok, tt.isValid)
		}
	}
}

func TestPhoneNumberNormalizer(t *testing.T) {
	tests := []struct {
		phoneNumber string
		newPrefix   string
		want        string
		wantErr     bool
	}{
		{"+989373708555", "0", "09373708555", false},
		{"989373708555", "0", "09373708555", false},
		{"00989022002580", "0", "09022002580", false},
		{"09122002580", "0", "09122002580", false},
		{"9322002580", "0", "09322002580", false},
		{"09373708555", "+98", "+989373708555", false},
		{"09022002580", "+98", "+989022002580", false},
		{"09122002580", "+98", "+989122002580", false},
		{"9322002580", "+98", "+989322002580", false},
		{"00989022002580", "+98", "+989022002580", false},
		{"09132222", "+98", "", true},
		{"9191282819921", "0", "", true},
	}

	for _, tt := range tests {
		got, err := PhoneNumberNormalizer(tt.phoneNumber, tt.newPrefix)
		if (err != nil) != tt.wantErr {
			t.Errorf("PhoneNumberNormalizer(%s, %s) error = %v, wantErr %t", tt.phoneNumber, tt.newPrefix, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("PhoneNumberNormalizer(%s, %s) = %s, want %s", tt.phoneNumber, tt.newPrefix, got, tt.want)
		}
	}
}

func TestGetOperatorPrefix(t *testing.T) {
	tests := []struct {
		phoneNumber string
		want        string
		wantErr     bool
	}{
		{"+989373708555", "937", false},
		{"00989013708555", "901", false},
		{"00988013708555", "", true},
	}

	for _, tt := range tests {
		got, err := GetOperatorPrefix(tt.phoneNumber)
		if (err != nil) != tt.wantErr {
			t.Errorf("GetOperatorPrefix(%s) error = %v, wantErr %t", tt.phoneNumber, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("GetOperatorPrefix(%s) = %s, want %s", tt.phoneNumber, got, tt.want)
		}
	}
}

func TestGetPhonePrefixOperator(t *testing.T) {
	expectedDetails904 := &OperatorDetails{
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Credit},
		operator: Irancell,
		model:    "سیم‌کارت کودک",
	}

	if details, err := GetPrefixDetails("904"); err != nil {
		t.Errorf("expected no error, got %v", err)
	} else if !reflect.DeepEqual(details, expectedDetails904) {
		t.Errorf("expected %v, got %v", expectedDetails904, details)
	}

	expectedDetails910 := &OperatorDetails{
		base:     "کشوری",
		province: []string{},
		simTypes: []SimType{Permanent, Credit},
		operator: MCI,
	}

	if details, err := GetPrefixDetails("910"); err != nil {
		t.Errorf("expected no error, got %v", err)
	} else if !reflect.DeepEqual(details, expectedDetails910) {
		t.Errorf("expected %v, got %v", expectedDetails910, details)
	}

	if _, err := GetPrefixDetails("9100"); err == nil {
		t.Error("expected error, got none")
	}
}

func TestGetPhoneDetails(t *testing.T) {
	expectedDetails := &OperatorDetails{
		base:     "تهران",
		province: []string{"البرز", "سمنان", "قم", "قزوین", "زنجان"},
		simTypes: []SimType{Credit},
		operator: MCI,
	}

	if details, err := GetPhoneDetails("09195431812"); err != nil {
		t.Errorf("expected no error, got %v", err)
	} else if !reflect.DeepEqual(details, expectedDetails) {
		t.Errorf("expected %v, got %v", expectedDetails, details)
	}

	if _, err := GetPhoneDetails("009195431812"); err == nil {
		t.Error("expected error, got none")
	}
}
