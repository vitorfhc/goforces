package scrapers

import "log"

type Target struct {
	Contest   string
	Problem   string
	IsContest bool
}

func Scrape(t Target) {
	if t.IsContest {
		log.Println("Scrapping contest", t.Contest)
		problems := ScrapeContest(t)

		for _, pt := range problems {
			ScrapeProblem(pt)
		}

		return
	}

	ScrapeProblem(t)
}

func ScrapeContest(t Target) []Target {
	return nil
}

func ScrapeProblem(t Target) {
	if t.IsContest {
		log.Fatal("Tchan")
	}

	log.Println("Scrapping problem", t.Problem, "from contest", t.Contest)
}
