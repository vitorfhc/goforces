package commands

import (
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test [binary]",
	Short: "Takes an executable and tests it with the input files",
	Long: `This command receives a binary files and by default
it tries to execute using all input-*.txt files as input
and compare with output=*.txt`,
	Args: cobra.ExactArgs(1),
	Run:  executeScraper,
}

func executeTest(cmd *cobra.Command, args []string) {}
