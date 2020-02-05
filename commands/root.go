package commands

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goforces",
	Short: "GoForces is a CodeForces scraper",
	Long: `To make a competitor's life easier GoForces
scrapes the problems inputs and outputs fastly.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
