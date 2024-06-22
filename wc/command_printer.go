package wc

import (
	"fmt"
)

func printAllCommandsOutput(
	printableCmd []PrintableCommand,
	bytes []byte,
	fileName string,
) {
	hasArgs := true
	for _, c := range printableCmd {
		c.Cmd.PrintOutput(&hasArgs, bytes, fileName)
	}
}

func printCalculationProcess(
	shouldPrint *bool,
	bytes []byte,
	fileName string,
	f func([]byte) int,
) {
	if !*shouldPrint {
		return
	}

	count := f(bytes)
	fmt.Printf("\t%d\t%s\n", count, fileName)

}

func PrintCommandOutput(
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
			printableCmd.HasArgs,
			bytes,
			fileName,
		)
	}

	if shouldDisplayAll {
		printAllCommandsOutput(
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
