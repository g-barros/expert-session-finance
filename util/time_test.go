package util

import "testing"

//TestStringToTime teste unitátio da função StringToTime
func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2020-01-01T23:10:11")

	if convertedTime.Year() != 2020 {
		t.Errorf("Espera que o ano seja 2020")
	}
}