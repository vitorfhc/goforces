package commands

import (
	"github.com/spf13/cobra"
	"github.com/vitorfhc/goforces/scrapers"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape [contest]",
	Short: "Scrapes a contest or a problem",
	Long: `Scrapes all the problems from a contest and saves them
in folders with its inputs and outputs.`,
	Args: cobra.ExactArgs(1),
	Run:  executeScraper,
}

func executeScraper(cmd *cobra.Command, args []string) {
	scrapers.Scrape(args[0])
}
