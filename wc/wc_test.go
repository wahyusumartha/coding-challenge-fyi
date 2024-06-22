package wc_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wahyusumartha/coding-challenges/wc"
	"testing"
)

func TestNumberOfBytes(t *testing.T) {
	runTest(func(bytes []byte) {
		// When
		b := wc.GetBytesLength(bytes)

		// Then
		assert.Equal(t, b, 12)
	})
}

func TestGetNumberOfLines(t *testing.T) {
	runTest(func(b []byte) {
		// When
		lines := wc.GetNumberOfLines(b)

		// Then
		assert.Equal(t, 2, lines)
	})
}

func TestGetNumberOfWords(t *testing.T) {
	runTest(func(b []byte) {
		// When
		words := wc.GetNumberOfWords(b)

		// Then
		assert.Equal(t, 2, words)
	})
}

func TestGetNumberOfCharacters(t *testing.T) {
	runTest(func(b []byte) {
		// When
		chars := wc.GetNumberOfCharacters(b)

		// Then
		assert.Equal(t, 12, chars)
	})
}

func runTest(sut func(b []byte)) {
	data := []byte("hello\nworld\n")
	sut(data)
}
