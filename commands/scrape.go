package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape [url]",
	Short: "Scrapes the problems from a url",
	Long: `Scrapes all the problems from a contest
or just the specific problem`,
	Args: cobra.ExactArgs(1),
	Run:  execute,
}

func execute(cmd *cobra.Command, args []string) {
	url := args[0]
	fmt.Println("Scrapping", url)
}
