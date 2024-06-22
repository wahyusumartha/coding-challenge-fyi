package wc

import (
	"strings"
	"unicode/utf8"
)

func GetBytesLength(b []byte) int {
	return len(b)
}

func GetNumberOfLines(b []byte) int {
	lines := strings.Split(string(b), "\n")
	if len(lines) > 0 {
		return len(lines) - 1
	} else {
		return 1
	}
}

func GetNumberOfWords(b []byte) int {
	words := strings.Fields(string(b))
	return len(words)
}

func GetNumberOfCharacters(b []byte) int {
	numberOfChar := utf8.RuneCount(b)
	return numberOfChar
}
