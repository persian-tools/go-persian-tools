package date

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jalaali/go-jalaali"
)

var ErrInvalidFormat error = fmt.Errorf("invalid date format")

func JalaaliToGregorianDate(jalaaliDate string) (time.Time, error) {
	var parts []string

	if strings.Contains(jalaaliDate, "/") {
		parts = strings.Split(jalaaliDate, "/")
	} else if strings.Contains(jalaaliDate, "-") {
		parts = strings.Split(jalaaliDate, "-")
	} else {
		return time.Time{}, ErrInvalidFormat
	}

	if len(parts) != 3 {
		return time.Time{}, ErrInvalidFormat
	}

	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return time.Time{}, err
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return time.Time{}, err
	}

	day, err := strconv.Atoi(parts[2])
	if err != nil {
		return time.Time{}, err
	}

	gYear, gMonth, gDay, err := jalaali.ToGregorian(year, jalaali.Month(month), day)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(gYear, time.Month(gMonth), gDay, 0, 0, 0, 0, time.UTC), nil
}

func JalaaliToGregorian(jalaaliDate string) (string, error) {
	var parts []string

	if strings.Contains(jalaaliDate, "/") {
		parts = strings.Split(jalaaliDate, "/")
	} else if strings.Contains(jalaaliDate, "-") {
		parts = strings.Split(jalaaliDate, "-")
	} else {
		return "", ErrInvalidFormat
	}

	if len(parts) != 3 {
		return "", ErrInvalidFormat
	}

	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", err
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}

	day, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", err
	}

	// Convert Jalaali to Gregorian
	gYear, gMonth, gDay, err := jalaali.ToGregorian(year, jalaali.Month(month), day)
	if err != nil {
		return "", err
	}
	// Return the Gregorian date as time.Time
	return fmt.Sprintf("%04d/%02d/%02d", gYear, gMonth, gDay), nil
}
