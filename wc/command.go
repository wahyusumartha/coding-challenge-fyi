package wc

import "io"

type Command interface {
	Usage() string
	Name() string
	Default() bool
	PrintOutput(w io.Writer, hasCommand *bool, bytes []byte, fileName string)
}

type BytesLengthCommand struct{}

var _ Command = (*BytesLengthCommand)(nil)

func (c BytesLengthCommand) Usage() string {
	return "The number of bytes in each input file is written to the standard output."
}

func (c BytesLengthCommand) Name() string {
	return "c"
}

func (c BytesLengthCommand) Default() bool {
	return false
}

func (c BytesLengthCommand) PrintOutput(w io.Writer, hasCommand *bool, bytes []byte, fileName string) {
	printCalculationProcess(
		w,
		hasCommand,
		bytes,
		fileName,
		GetBytesLength,
	)
}

type LineCounterCommand struct{}

var _ Command = (*LineCounterCommand)(nil)

func (l LineCounterCommand) Usage() string {
	return "The number of lines in each input file is written to the standard output."
}

func (l LineCounterCommand) Name() string {
	return "l"
}

func (l LineCounterCommand) Default() bool {
	return false
}

func (l LineCounterCommand) PrintOutput(w io.Writer, hasCommand *bool, bytes []byte, fileName string) {
	printCalculationProcess(
		w,
		hasCommand,
		bytes,
		fileName,
		GetNumberOfLines,
	)
}

type WordCounterCommand struct{}

var _ Command = (*WordCounterCommand)(nil)

func (w WordCounterCommand) Usage() string {
	return "The number of words in each input file is written to the standard output."
}

func (w WordCounterCommand) Name() string {
	return "w"
}

func (w WordCounterCommand) Default() bool {
	return false
}

func (w WordCounterCommand) PrintOutput(writer io.Writer, hasCommand *bool, bytes []byte, fileName string) {
	printCalculationProcess(
		writer,
		hasCommand,
		bytes,
		fileName,
		GetNumberOfWords,
	)
}

type CharacterCounterCommand struct{}

var _ Command = (*CharacterCounterCommand)(nil)

func (c CharacterCounterCommand) Usage() string {
	return "The number of characters in each input file is written to the standard output.\nIf the current locale does not\nsupport multibyte characters, this is equivalent to the -c option."
}

func (c CharacterCounterCommand) Name() string {
	return "m"
}

func (c CharacterCounterCommand) Default() bool {
	return false
}

func (c CharacterCounterCommand) PrintOutput(w io.Writer, hasCommand *bool, bytes []byte, fileName string) {
	printCalculationProcess(
		w,
		hasCommand,
		bytes,
		fileName,
		GetNumberOfCharacters,
	)
}
