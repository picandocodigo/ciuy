package ciuy_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/picandocodigo/ciuy"
)

func TestCiUyFunc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CiUy Suite")
}

var _ = Describe("Validation Digit", func() {
	Context("given an initial ci", func() {
		It("returns the validation digit for random ci", func() {
			digit := ciuy.ValidationDigit("6098714")
			Expect(digit).To(Equal("9"))
		})
		It("returns the validation digit for another random ci", func() {
			digit := ciuy.ValidationDigit("9575350")
			Expect(digit).To(Equal("3"))
		})
		It("returns the validation digit for 1", func() {
			digit := ciuy.ValidationDigit("1111111")
			Expect(digit).To(Equal("1"))
		})
		It("returns the validation digit for 2", func() {
			digit := ciuy.ValidationDigit("2222222")
			Expect(digit).To(Equal("2"))
		})
	})
})

var _ = Describe("Transform Ci", func() {
	Context("given a dirty string", func() {
		It("returns just the digits for dots", func() {
			ci := ciuy.Transform("1.111.111-1")
			expected := "11111111"
			Expect(ci).To(Equal(expected))
		})
		It("returns just the digits for dots and dash", func() {
			ci := ciuy.Transform("2.222.222-2")
			expected := "22222222"
			Expect(ci).To(Equal(expected))
		})
		It("returns just the digits for dots and slash", func() {
			ci := ciuy.Transform("3.333.333/3")
			expected := "33333333"
			Expect(ci).To(Equal(expected))
		})
	})
})

var _ = Describe("Validate Ci", func() {
	Context("Correct Ci numbers", func() {
		It("validates a right CI", func() {
			ci := "1.111.111-1"
			result := ciuy.ValidateCi(ci)
			Expect(result).To(Equal(true))
		})
		It("validates a random right CI", func() {
			ci := "9.575.350/3"
			result := ciuy.ValidateCi(ci)
			Expect(result).To(Equal(true))
		})
		It("validates cis with 6 digits", func() {
			ci := "111,111_3"
			result := ciuy.ValidateCi(ci)
			Expect(result).To(Equal(true))
		})
	})
	Context("Wrong Ci numbers are not valid", func() {
		It("Doesn't validate a wrong Ci", func() {
			ci := "1.111.111-4"
			result := ciuy.ValidateCi(ci)
			Expect(result).To(Equal(false))
		})
		It("Does not validate a wrong random ci", func() {
			ci := "9.575.350/8"
			result := ciuy.ValidateCi(ci)
			Expect(result).To(Equal(false))
		})
	})
})
