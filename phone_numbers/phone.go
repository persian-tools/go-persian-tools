package phonenumbers

import (
	"errors"
	"strings"
)

// GetPhoneDetails returns the OperatorDetails for the given phoneNumber
func GetPhoneDetails(phoneNumber string) (*OperatorDetails, error) {
	if ok := IsPhoneValid(phoneNumber); !ok {
		return nil, ErrInvalidFormat
	}
	prefix, err := GetOperatorPrefix(phoneNumber)
	if err != nil {
		return nil, err
	}
	return GetPrefixDetails(prefix)
}

func GetPrefixDetails(prefix string) (*OperatorDetails, error) {
	// Define a slice of maps to iterate over
	operatorsMap := []map[string]OperatorDetails{
		MCIMap,
		TALIYA,
		RIGHTTEL,
		IRANCELL,
		SHATELMOBILE,
	}

	// Iterate over each map and check for the prefix
	for _, m := range operatorsMap {
		if details, found := m[prefix]; found {
			return &details, nil
		}
	}

	return nil, errors.New("Invalid prefix")
}

// IsPhoneValid returns if the phoneNumber is valid
func IsPhoneValid(phoneNumber string) bool {
	prefix := GetPhonePrefix(phoneNumber)

	phoneNumberWithoutPrefix := phoneNumber[len(prefix):]
	if len(phoneNumberWithoutPrefix) == 10 && phoneNumberWithoutPrefix[0] == '9' {
		return true
	}

	return false
}

// GetPhonePrefix returns the prefix for Iranian phone numbers
// Example +98 98 0098 0
func GetPhonePrefix(phoneNumber string) string {
	for _, prefix := range Prefixes {
		if strings.HasPrefix(phoneNumber, prefix) {
			return prefix
		}
	}
	return ""
}

// phoneNumberNormalizer replaces the current phone number prefix with the desired prefix
func PhoneNumberNormalizer(phoneNumber, newPrefix string) (string, error) {
	if ok := IsPhoneValid(phoneNumber); !ok {
		return "", ErrInvalidFormat
	}

	prefix := GetPhonePrefix(phoneNumber)

	return newPrefix + phoneNumber[len(prefix):], nil
}

// GetOperatorPrefix returns the operator prefix of the phone number
func GetOperatorPrefix(phoneNumber string) (string, error) {
	if ok := IsPhoneValid(phoneNumber); !ok {
		return "", ErrInvalidFormat
	}

	for _, prefix := range Prefixes {
		if strings.HasPrefix(phoneNumber, prefix) {
			return phoneNumber[len(prefix) : len(prefix)+3], nil
		}
	}

	return "", ErrInvalidFormat
}
