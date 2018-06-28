package ciuy

import (
	"regexp"
	"strconv"
)

// ValidationDigit returns the validation digit given a string for a number
func ValidationDigit(blob string) string {
	ci := Transform(blob)
	if len(ci) == 6 {
		ci = "0" + ci
	}
	a := 0
	validationAlg := "2987634"
	var validationDigit, ciDigit int
	var err error
	for index, digit := range validationAlg {
		validationDigit, err = strconv.Atoi(string(digit))
		if err != nil {
			panic("FML!")
		}
		ciDigit, err = strconv.Atoi(string(ci[index]))
		if err != nil {
			panic("FML!")
		}
		a += validationDigit * ciDigit
	}
	validationDigit = (10 - (a % 10))
	digitString := strconv.Itoa(validationDigit)
	return digitString
}

// Transform receives a string with mixed characters returns the digits as a string
func Transform(ci string) string {
	re := regexp.MustCompile(`[^\d]`)
	cleanCi := re.ReplaceAllString(ci, "")
	return cleanCi
}
