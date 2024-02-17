package util

import "testing"

func TestStringToDateTimeShouldSuccess(t *testing.T) {
	var convertedTime = StringToDateTime("2024-02-29T15:45:00")

	if convertedTime.Year() != 2024 {
		t.Errorf("Expect year 2024")
	}

	if convertedTime.Month() != 02 {
		t.Errorf("Expect month to be 02")
	}

	if convertedTime.Day() != 29 {
		t.Errorf("Expect day to be 29")
	}

	if convertedTime.Hour() != 15 {
		t.Errorf("Expect month to be 15")
	}

	if convertedTime.Minute() != 45 {
		t.Errorf("Expect month to be 45")
	}

	if convertedTime.Second() != 00 {
		t.Errorf("Expect month to be 00")
	}
}
