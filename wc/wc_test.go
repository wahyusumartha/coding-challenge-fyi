package wc_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wahyusumartha/coding-challenges/wc"
	"os"
	"testing"
)

func TestNumberOfBytes(t *testing.T) {
	runTest(func(fileName string) {
		// When
		b, err := wc.GetBytes(fileName)

		// Then
		assert.Equal(t, b, 12)
		assert.Nil(t, err)
	})
}

func TestNumberOfBytes_ShouldReturnError_WhenFailedReadingFile(t *testing.T) {
	runTest(func(string) {
		// When
		b, err := wc.GetBytes("not_found.txt")

		// Then
		assert.Equal(t, b, 0)
		assert.NotNil(t, err)
	})
}

func TestGetNumberOfLines(t *testing.T) {
	runTest(func(fileName string) {
		// When
		lines, err := wc.GetNumberOfLines(fileName)

		// Then
		assert.Nil(t, err)
		assert.Equal(t, 2, lines)
	})
}

func TestGetNumberOfLines_ShouldReturnError_WhenFailedReadingFile(t *testing.T) {
	runTest(func(string) {
		// When
		lines, err := wc.GetNumberOfLines("not_found.txt")

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, 0, lines)
	})
}

func TestGetNumberOfWords(t *testing.T) {
	runTest(func(fileName string) {
		// When
		words, err := wc.GetNumberOfWords(fileName)

		// Then
		assert.Nil(t, err)
		assert.Equal(t, 2, words)
	})
}

func TestGetNumberOfWords_ShouldReturnError_WhenFailedReadingFile(t *testing.T) {
	runTest(func(string) {
		// When
		words, err := wc.GetNumberOfWords("not_found.txt")

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, 0, words)
	})
}

func TestGetNumberOfCharacters(t *testing.T) {
	runTest(func(fileName string) {
		// When
		chars, err := wc.GetNumberOfCharacters(fileName)

		// Then
		assert.Nil(t, err)
		assert.Equal(t, 12, chars)
	})
}

func TestGetNumberOfCharacters_ShouldReturnError_WhenFailedReadingFile(t *testing.T) {
	runTest(func(string) {
		// When
		chars, err := wc.GetNumberOfCharacters("not_found.txt")

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, 0, chars)
	})
}

func prepareFile(name string) ([]byte, error) {
	data := []byte("hello\nworld\n")
	err := os.WriteFile(name, data, 0644)

	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func runTest(sut func(filename string)) {
	fileName := "/tmp/data-xxx"
	_, err := prepareFile(fileName)
	if err != nil {
		panic(err)
	}

	sut(fileName)

	err = os.Remove(fileName)
	if err != nil {
		panic(err)
	}
}
