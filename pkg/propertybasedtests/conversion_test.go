package romannumaerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

type ConvertToRomanTestCase struct {
	Arabic uint32
	Roman  string
}

var cases = []ConvertToRomanTestCase{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %s", test.Arabic, test.Roman), func(t *testing.T) {
			expected := test.Roman
			actual := ConvertToRoman(test.Arabic)

			if actual != expected {
				t.Errorf("Actual: %q - Expected: %q", actual, expected)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("%s gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			expected := test.Arabic
			actual := ConvertToArabic(test.Roman)

			if actual != expected {
				t.Errorf("Actual: %d - Expected: %d", actual, expected)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint32) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		romanFromArabic := ConvertToRoman(arabic)
		arabicFromRoman := ConvertToArabic(romanFromArabic)
		return arabic == arabicFromRoman
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 5000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
