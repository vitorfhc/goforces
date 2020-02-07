package scrapers

import (
	"log"
	"os"

	"github.com/gocolly/colly"
)

var (
	collector *colly.Collector
	contest   string
)

const (
	urlPrefix   = "https://codeforces.com"
	contestPath = "/contest/"
)

func init() {
	collector = colly.NewCollector(
		colly.AllowedDomains("www.codeforces.com", "codeforces.com"),
	)

	collector.OnHTML("table.problems tbody", ScrapeContest)
}

func Scrape(contestNumber string) {
	contest = contestNumber
	contestURL := urlPrefix + contestPath + contest
	err := collector.Visit(contestURL)
	if err != nil {
		log.Fatal(err)
	}
}

func ScrapeContest(el *colly.HTMLElement) {
	err := os.MkdirAll(contest, 0644)
	if err != nil {
		log.Fatal(err)
	}

	el.ForEach("tr td:nth-of-type(1) a", func(_ int, e *colly.HTMLElement) {
		collector.Visit(urlPrefix + e.Attr("href"))
	})
}

func ScrapeProblem() {}
