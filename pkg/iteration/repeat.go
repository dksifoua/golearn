package iteration

import "strings"

func Repeat(character string, count int) (repeated string) {
	for i := 0; i < count; i++ {
		repeated += character
	}

	return
}

func RepeatOptimal(character byte, count int) string {
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteByte(character)
	}

	return repeated.String()
}
