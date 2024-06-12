package scrape

import (
	"errors"
	"fmt"
	"io"
	"time"
)

import (
	goq "github.com/PuerkitoBio/goquery"
)

type Submission struct {
	score       float64
	participant string
	team        string
	timestamp   time.Time
}

func scrapeTask(r io.Reader) ([]Submission, error) {
	doc, err := goq.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	var subms []Submission
	doc.Find("div#leaderboard table tbody tr").EachWithBreak(func(_ int, s *goq.Selection) bool {
		var subm Submission
		subm, err = scrapeSubmission(s)
		if err != nil {
			return false
		}
		subms = append(subms, subm)
		return true
	})
	if err != nil {
		return nil, err
	}
	return subms, nil
}

func scrapeSubmission(s *goq.Selection) (Submission, error) {
	var subm Submission
	cells := eachToStr(s.Find("td"))
	if len(cells) < 4 {
		return subm, errors.New("Incomplete submission data")
	}

	timestamp, err := time.ParseInLocation("02.01.2006 15:04", cells[2], time.UTC)
	if err != nil {
		return subm, err
	}

	_, err = fmt.Sscan(cells[0], &subm.score)
	if err != nil {
		return subm, err
	}
	subm.participant = cells[1]
	subm.timestamp = timestamp
	subm.team = cells[3]
	return subm, nil
}

func eachToStr(s *goq.Selection) []string {
	strs := make([]string, s.Length())
	s.Each(func(i int, s *goq.Selection) {
		strs[i] = s.Text()
	})
	return strs
}
