package commands

import (
	"log"

	"github.com/spf13/cobra"
)

type target struct {
	contest   string
	problem   string
	isContest bool
}

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
	t := target{"", "", false}
	t.contest = args[0]

	if len(args) == 1 {
		t.isContest = true
		scrape(t)
		return
	}

	t.problem = args[1]
	scrape(t)
}

func scrape(t target) {
	if t.isContest {
		log.Println("Scrapping contest", t.contest)
		problems := scrapeContest(t)

		for _, pt := range problems {
			scrapeProblem(pt)
		}

		return
	}

	scrapeProblem(t)
}

func scrapeContest(t target) []target {
	return nil
}

func scrapeProblem(t target) {
	log.Println("Scrapping problem", t.problem, "from contest", t.contest)
}
