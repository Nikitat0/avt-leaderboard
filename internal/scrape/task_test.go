package scrape

import (
	_ "embed"
	"strings"
	"testing"
	"time"
)

import (
	"github.com/stretchr/testify/assert"
)

//go:embed testdata/task.html
var taskHtml string

func TestScrapeTask(t *testing.T) {
	var expected = []Submission{
		{
			score:       0.8,
			participant: "Alice",
			timestamp:   time.Date(2024, time.January, 13, 12, 30, 0, 0, time.UTC),
			team:        "TeamFoo",
		},
		{
			score:       0.6,
			participant: "Bob",
			timestamp:   time.Date(2024, time.February, 14, 13, 40, 0, 0, time.UTC),
			team:        "TeamBar",
		},
		{
			score:       0.4,
			participant: "Mallory",
			timestamp:   time.Date(2024, time.March, 15, 14, 50, 0, 0, time.UTC),
			team:        "",
		},
	}

	r := strings.NewReader(taskHtml)
	got, err := scrapeTask(r)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, got)
}
