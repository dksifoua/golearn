package romannumaerals

import "strings"

type Numeral struct {
	Arabic uint32
	Roman  string
}

var numerals = []Numeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint32) string {
	var roman strings.Builder
	for _, numeral := range numerals {
		for arabic >= numeral.Arabic {
			roman.WriteString(numeral.Roman)
			arabic -= numeral.Arabic
		}
	}

	return roman.String()
}

func ConvertToArabic(roman string) uint32 {
	var arabic uint32
	for _, numeral := range numerals {
		for strings.HasPrefix(roman, numeral.Roman) {
			arabic += numeral.Arabic
			roman = strings.TrimPrefix(roman, numeral.Roman)
		}
	}

	return arabic
}
