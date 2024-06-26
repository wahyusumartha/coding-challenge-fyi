package wc

import (
	"fmt"
	"io"
)

func printAllCommandsOutput(
	w io.Writer,
	printableCmd []PrintableCommand,
	bytes []byte,
	fileName string,
) {
	hasArgs := true
	for _, c := range printableCmd {
		c.Cmd.PrintOutput(w, &hasArgs, bytes, fileName)
	}
}

func printCalculationProcess(
	outputWriter io.Writer,
	shouldPrint *bool,
	bytes []byte,
	fileName string,
	f func([]byte) int,
) {
	if !*shouldPrint {
		return
	}

	count := f(bytes)
	_, _ = fmt.Fprintf(outputWriter, "\t%d\t%s\n", count, fileName)
}

func PrintCommandOutput(
	w io.Writer,
	bytes []byte,
	fileName string,
	printableCmds []PrintableCommand,
) {
	shouldDisplayAll := true
	for _, printableCmd := range printableCmds {
		if *printableCmd.HasArgs {
			shouldDisplayAll = false
		}
		printableCmd.Cmd.PrintOutput(
			w,
			printableCmd.HasArgs,
			bytes,
			fileName,
		)
	}

	if shouldDisplayAll {
		printAllCommandsOutput(
			w,
			printableCmds,
			bytes,
			fileName,
		)
	}
}

type PrintableCommand struct {
	Cmd     Command
	HasArgs *bool
}
