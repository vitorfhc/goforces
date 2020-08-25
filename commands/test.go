package commands

import (
	"github.com/spf13/cobra"
	"github.com/vitorfhc/goforces/tester"
)

var testCmd = &cobra.Command{
	Use:   "test [binary]",
	Short: "Takes an executable and tests it with the input files",
	Long: `This command receives an executable and makes tests
with all input files comparing then with the output files.
Tip: you can run python files using the shebang #!/bin/python`,
	Args: cobra.ExactArgs(1),
	Run:  executeTest,
}

func executeTest(cmd *cobra.Command, args []string) {
	tester.Test(args[0])
}
