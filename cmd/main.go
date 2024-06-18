package main

import (
	"flag"
	"fmt"
	"github.com/wahyusumartha/coding-challenges/wc"
	"os"
)

func main() {
	var c, l, w, m string
	const cOptionsUsage = "The number of bytes in each input file is written to the standard output."
	const lOptionsUsage = "The number of lines in each input file is written to the standard output."
	const wOptionsUsage = "The number of words in each input file is written to the standard output."
	const mOptionsUsage = "The number of characters in each input file is written to the standard output. \nIf the current locale does not support multibyte characters, this is equivalent to the -c option.  This will"
	flag.StringVar(&c, "c", "", cOptionsUsage)
	flag.StringVar(&l, "l", "", lOptionsUsage)
	flag.StringVar(&w, "w", "", wOptionsUsage)
	flag.StringVar(&m, "m", "", mOptionsUsage)
	flag.Parse()

	if len(c) == 0 && len(l) == 0 && len(w) == 0 && len(m) == 0 {
		arg := os.Args[1]
		count, err := wc.GetBytes(arg)
		lines, err := wc.GetNumberOfLines(arg)
		words, err := wc.GetNumberOfWords(arg)

		if err != nil {
			panic(err)
		}
		fmt.Printf("%d %d %d %s\n", lines, words, count, arg)
	}

	if len(c) > 0 {
		count, err := wc.GetBytes(c)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d %s\n", count, c)
	}

	if len(l) > 0 {
		lines, err := wc.GetNumberOfLines(l)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d %s\n", lines, l)
	}

	if len(w) > 0 {
		words, err := wc.GetNumberOfWords(w)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d %s\n", words, w)

	}

	if len(m) > 0 {
		chars, err := wc.GetNumberOfCharacters(m)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d %s\n", chars, m)
	}

}
