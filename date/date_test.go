package date

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestJalaaliToGregorianDate(t *testing.T) {
	tests := []struct {
		input       string
		outPut      time.Time
		err         error
		shouldError bool
	}{
		{"1401/02/03", time.Date(2022, 4, 23, 0, 0, 0, 0, time.UTC), nil, false},
		{"1401-02-03", time.Date(2022, 4, 23, 0, 0, 0, 0, time.UTC), nil, false},
		{"1401-02/03", time.Time{}, ErrInvalidFormat, true},
		{"1401/02-03", time.Time{}, ErrInvalidFormat, true},
		{"1401-02-0A", time.Time{}, nil, true},
		{"1401-0B-03", time.Time{}, nil, true},
		{"140C-02-03", time.Time{}, nil, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing input: %s", tt.input), func(t *testing.T) {
			georgianDate, err := JalaaliToGregorianDate(tt.input)
			if tt.shouldError {
				if err == nil {
					t.Errorf("expected error for input %v but got none", tt.input)
				} else if tt.err != nil && !errors.Is(err, tt.err) {
					t.Errorf("unexpected error for input %v: got %v, expected %v", tt.input, err, tt.err)
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error for input %v but got %v", tt.input, err)
				} else if !georgianDate.Equal(tt.outPut) {
					t.Errorf("wrong result for input %v: got %v, expected %v", tt.input, georgianDate, tt.outPut)
				}
			}
		})
	}
}

func TestJalaaliToGregorian(t *testing.T) {
	tests := []struct {
		input       string
		outPut      string
		err         error
		shouldError bool
	}{
		{"1401/02/03", "2022/04/23", nil, false},
		{"1401-02-03", "2022/04/23", nil, false},
		{"1401-02/03", "", ErrInvalidFormat, true},
		{"1401/02-03", "", ErrInvalidFormat, true},
		{"1401-02-0A", "", nil, true},
		{"1401-0B-03", "", nil, true},
		{"140C-02-03", "", nil, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing input: %s", tt.input), func(t *testing.T) {
			georgian, err := JalaaliToGregorian(tt.input)
			if tt.shouldError {
				if err == nil {
					t.Errorf("expected error for input %v but got none", tt.input)
				} else if tt.err != nil && !errors.Is(err, tt.err) {
					t.Errorf("unexpected error for input %v: got %v, expected %v", tt.input, err, tt.err)
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error for input %v but got %v", tt.input, err)
				} else if georgian != tt.outPut {
					t.Errorf("wrong result for input %v: got %v, expected %v", tt.input, georgian, tt.outPut)
				}
			}
		})
	}

}
