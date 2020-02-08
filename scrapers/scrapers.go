package scrapers

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var (
	collector *colly.Collector
	contest   string
	workdir   string
)

const (
	urlPrefix   = "https://codeforces.com"
	contestPath = "/contest/"
)

func init() {
	collector = colly.NewCollector(
		colly.AllowedDomains("www.codeforces.com", "codeforces.com"),
	)

	collector.OnHTML("table.problems tbody", scrapeContest)
	collector.OnHTML("div.problem-statement", scrapeProblem)

	var err error
	workdir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	workdir = workdir + "/"
}

// Scrape given contest where contest number is the number in the url
func Scrape(contestNumber string) {
	contest = contestNumber
	contestURL := urlPrefix + contestPath + contest
	err := collector.Visit(contestURL)
	if err != nil {
		log.Fatal(err)
	}
}

func scrapeContest(el *colly.HTMLElement) {
	path := workdir + contest
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	el.ForEach("tr td:nth-of-type(1) a", func(_ int, e *colly.HTMLElement) {
		collector.Visit(urlPrefix + e.Attr("href"))
	})
}

func scrapeProblem(el *colly.HTMLElement) {
	problemTitle := el.ChildText("div.title")
	problemShortName := problemTitle[:strings.IndexByte(problemTitle, '.')]
	problemFolder := workdir + contest + "/" + problemShortName + "/"
	err := os.Mkdir(problemFolder, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	el.ForEach("div.sample-test", func(_ int, e *colly.HTMLElement) {
		testNumber := 1
		e.ForEach("pre", func(ind int, pre *colly.HTMLElement) {
			isInput := ind%2 == 0

			filename := buildProblemFilename(testNumber, isInput)
			path := problemFolder + filename
			writeTestFile(path, pre.Text)

			if !isInput {
				testNumber++
			}
		})
	})
}

func buildProblemFilename(ind int, isInput bool) string {
	var filename string

	if isInput {
		filename = "input-" + strconv.Itoa(ind) + ".txt"
	} else {
		filename = "output-" + strconv.Itoa(ind) + ".txt"
	}

	return filename
}

func writeTestFile(path string, text string) {
	f, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.WriteString(text)
}
