package main

import (
	"flag"
	"fmt"
	"github.com/wahyusumartha/coding-challenges/wc"
	"io"
	"os"
)

func main() {
	commands := []wc.Command{
		wc.BytesLengthCommand{},
		wc.LineCounterCommand{},
		wc.WordCounterCommand{},
		wc.CharacterCounterCommand{},
	}
	outputCmd := parseCommands(commands)
	input, err := determineInput(flag.Args())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "[ccwc] %s\n", err)
		os.Exit(1)
	}

	b := input.byte
	fileName := input.fileName

	wc.PrintCommandOutput(b, fileName, outputCmd)
	os.Exit(0)
}

func parseCommands(commands []wc.Command) []wc.PrintableCommand {
	var cmd []wc.PrintableCommand
	for _, c := range commands {
		f := flag.Bool(c.Name(), c.Default(), c.Usage())
		cmd = append(cmd, wc.PrintableCommand{Cmd: c, HasArgs: f})
	}

	flag.Parse()
	return cmd
}

type input struct {
	byte     []byte
	fileName string
}

func determineInput(remainingArgs []string) (input, error) {
	hasFileName := len(remainingArgs) > 0
	if hasFileName {
		fileName := remainingArgs[0]
		return readInputFromFile(fileName)
	}

	return readInputFromStdin()
}

func readInputFromFile(fileName string) (input, error) {
	b, err := os.ReadFile(fileName)
	input := input{
		b,
		fileName,
	}
	return input, err
}

func readInputFromStdin() (input, error) {
	b, err := io.ReadAll(os.Stdin)
	input := input{
		b,
		"",
	}
	return input, err
}
