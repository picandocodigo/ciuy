package ciuy

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
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
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(ci, "")
}

// ValidateCi gets a Ci string and returns a bool
func ValidateCi(ci string) bool {
	ci = Transform(ci)
	if len(ci) < 6 {
		return false
	}
	dig := string(ci[len(ci)-1])
	ci = ci[0 : len(ci)-1]
	return ValidationDigit(ci) == dig
}

// Random creates a random valid Ci number
func Random() string {
	rand.Seed(time.Now().Unix())
	max := 9999999
	min := 1000000
	ci := strconv.Itoa(rand.Intn(max-min) + min)
	result := ci + ValidationDigit(ci)
	return result
}
