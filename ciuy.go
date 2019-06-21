package ciuy

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// ValidationDigit returns the validation digit given a string for a number
func ValidationDigit(blob string) (string, error) {
	ci := Transform(blob)
	if len(ci) == 6 {
		ci = "0" + ci
	}
	a := 0
	validationAlg := []int{2, 9, 8, 7, 6, 3, 4}
	for index, digit := range validationAlg {
		ciDigit, err := strconv.Atoi(string(ci[index]))
		if err != nil {
			return "", err
		}
		a += digit * ciDigit
	}
	if mod := a % 10; mod != 0 {
		return strconv.Itoa(10 - mod), nil
	}
	return "0", nil
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
	digit, err := ValidationDigit(ci)
	return err == nil && digit == dig
}

// Random creates a random valid Ci number
func Random() string {
	rand.Seed(time.Now().Unix())
	max := 9999999
	min := 1000000
	ci := strconv.Itoa(rand.Intn(max-min) + min)
	digit, err := ValidationDigit(ci)
	if err != nil {
		panic("FML!")
	}
	return ci + digit
}
