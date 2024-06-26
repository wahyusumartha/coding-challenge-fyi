package wc_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/wahyusumartha/coding-challenges/wc"
	"io"
	"testing"
)

type input struct {
	commands []wc.PrintableCommand
	bytes    []byte
	fileName string
}

type tableTest struct {
	name     string
	input    input
	expected string
}

func TestPrintCommandOutput_WhenInputStdIn(t *testing.T) {
	data := []byte("hello\nworld\n")
	hasArgs := true
	hasNoArgs := false
	tests := []tableTest{
		{
			name: "count bytes length",
			input: input{
				commands: []wc.PrintableCommand{
					{
						wc.BytesLengthCommand{},
						&hasArgs,
					},
				},
				bytes: data,
			},
			expected: "\t12\t\n",
		},
		{
			name: "count number of lines",
			input: input{
				commands: []wc.PrintableCommand{
					{
						wc.LineCounterCommand{},
						&hasArgs,
					},
				},
				bytes: data,
			},
			expected: "\t2\t\n",
		},
		{
			name: "count number of words",
			input: input{
				commands: []wc.PrintableCommand{
					{
						wc.WordCounterCommand{},
						&hasArgs,
					},
				},
				bytes: data,
			},
			expected: "\t2\t\n",
		},
		{
			name: "count number of character",
			input: input{
				commands: []wc.PrintableCommand{
					{
						wc.CharacterCounterCommand{},
						&hasArgs,
					},
				},
				bytes: data,
			},
			expected: "\t12\t\n",
		},
		{
			name: "arguments not specified should show all",
			input: input{
				commands: []wc.PrintableCommand{
					{
						wc.BytesLengthCommand{},
						&hasNoArgs,
					},
					{
						wc.LineCounterCommand{},
						&hasNoArgs,
					},
					{
						wc.WordCounterCommand{},
						&hasNoArgs,
					},
					{
						wc.CharacterCounterCommand{},
						&hasNoArgs,
					},
				},
				bytes: data,
			},
			expected: "\t12\t\n\t2\t\n\t2\t\n\t12\t\n",
		},
		{
			name: "some arguments specified should only show line and characters",
			input: input{
				commands: []wc.PrintableCommand{
					{
						wc.BytesLengthCommand{},
						&hasNoArgs,
					},
					{
						wc.LineCounterCommand{},
						&hasArgs,
					},
					{
						wc.WordCounterCommand{},
						&hasNoArgs,
					},
					{
						wc.CharacterCounterCommand{},
						&hasArgs,
					},
				},
				bytes: data,
			},
			expected: "\t2\t\n\t12\t\n",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				test.expected,
				captureStdout(
					test.input.bytes,
					test.input.fileName,
					test.input.commands,
					wc.PrintCommandOutput,
				),
			)
		})
	}
}

func TestPrintCommandOutput_WhenInputFileName(t *testing.T) {
	data := []byte("hello\nworld\n")
	hasArgs := true
	hasNoArgs := false
	fileName := "test.txt"
	tests := []tableTest{
		{
			name: "count bytes length",
			input: input{
				commands: []wc.PrintableCommand{
					{
						wc.BytesLengthCommand{},
						&hasArgs,
					},
				},
				bytes:    data,
				fileName: fileName,
			},
			expected: "\t12\t" + fileName + "\n",
		},
		{
			name: "count number of lines",
			input: input{
				commands: []wc.PrintableCommand{
					{
						wc.LineCounterCommand{},
						&hasArgs,
					},
				},
				bytes:    data,
				fileName: fileName,
			},
			expected: "\t2\t" + fileName + "\n",
		},
		{
			name: "count number of words",
			input: input{
				commands: []wc.PrintableCommand{
					{
						wc.WordCounterCommand{},
						&hasArgs,
					},
				},
				bytes:    data,
				fileName: fileName,
			},
			expected: "\t2\t" + fileName + "\n",
		},
		{
			name: "count number of character",
			input: input{
				commands: []wc.PrintableCommand{
					{
						wc.CharacterCounterCommand{},
						&hasArgs,
					},
				},
				bytes:    data,
				fileName: fileName,
			},
			expected: "\t12\t" + fileName + "\n",
		},
		{
			name: "arguments not specified should show all",
			input: input{
				commands: []wc.PrintableCommand{
					{
						wc.BytesLengthCommand{},
						&hasNoArgs,
					},
					{
						wc.LineCounterCommand{},
						&hasNoArgs,
					},
					{
						wc.WordCounterCommand{},
						&hasNoArgs,
					},
					{
						wc.CharacterCounterCommand{},
						&hasNoArgs,
					},
				},
				bytes:    data,
				fileName: fileName,
			},
			expected: "\t12\t" + fileName + "\n\t2\t" + fileName + "\n\t2\t" + fileName + "\n\t12\t" + fileName + "\n",
		},
		{
			name: "some arguments specified should only show line and word",
			input: input{
				commands: []wc.PrintableCommand{
					{
						wc.BytesLengthCommand{},
						&hasNoArgs,
					},
					{
						wc.LineCounterCommand{},
						&hasArgs,
					},
					{
						wc.WordCounterCommand{},
						&hasArgs,
					},
					{
						wc.CharacterCounterCommand{},
						&hasNoArgs,
					},
				},
				bytes:    data,
				fileName: fileName,
			},
			expected: "\t2\t" + fileName + "\n\t2\t" + fileName + "\n",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				test.expected,
				captureStdout(
					test.input.bytes,
					test.input.fileName,
					test.input.commands,
					wc.PrintCommandOutput,
				),
			)
		})
	}
}

func captureStdout(
	b []byte,
	fileName string,
	cmd []wc.PrintableCommand,
	f func(io.Writer, []byte, string, []wc.PrintableCommand),
) string {
	var buf bytes.Buffer
	f(&buf, b, fileName, cmd)
	return buf.String()
}
