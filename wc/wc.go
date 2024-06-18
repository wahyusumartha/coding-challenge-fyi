package wc

import (
	"os"
	"strings"
	"unicode/utf8"
)

func GetBytes(fileName string) (int, error) {
	b, err := readFile(fileName)

	if err != nil {
		return 0, err
	}

	return len(b), nil
}

func GetNumberOfLines(fileName string) (int, error) {
	b, err := readFile(fileName)

	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(b), "\n")
	if len(lines) > 0 {
		return len(lines) - 1, nil
	} else {
		return 1, nil
	}
}

func GetNumberOfWords(fileName string) (int, error) {
	b, err := readFile(fileName)

	if err != nil {
		return 0, err
	}

	words := strings.Fields(string(b))
	return len(words), nil
}

func GetNumberOfCharacters(fileName string) (int, error) {
	b, err := readFile(fileName)

	if err != nil {
		return 0, err
	}

	numberOfChar := utf8.RuneCount(b)
	return numberOfChar, nil
}

func readFile(filename string) ([]byte, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}
