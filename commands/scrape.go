package commands

import (
	"github.com/spf13/cobra"
	"github.com/vitorfhc/goforces/scrapers"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape [contest] [problem]",
	Short: "Scrapes a contest or a problem",
	Long: `Scrapes all the problems from a contest if no problem is given
or just the specific problem.
Contest: saves each problem in a specific folder
Problem: saves every inputs, outputs, and the problem's text`,
	Args: cobra.RangeArgs(1, 2),
	Run:  execute,
}

func execute(cmd *cobra.Command, args []string) {
	t := scrapers.Target{"", "", false}
	t.Contest = args[0]

	if len(args) == 1 {
		t.IsContest = true
		scrapers.Scrape(t)
		return
	}

	t.Problem = args[1]
	scrapers.Scrape(t)
}
