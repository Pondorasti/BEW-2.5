package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type Jobs struct {
	Title     string `json:"title"`
	Link      string `json:"link"`
	Timestamp string `json:"timestamp"`
}

func normalizeLink(link string) string {
	if link[0] == 'i' { // item?id=
		link = "https://news.ycombinator.com/" + link
	}

	return link
}

func main() {
	jobs := []Jobs{}
	timestampIndex := 0
	c := colly.NewCollector()

	// Job Details
	c.OnHTML("tr.athing > td.title > a.titlelink", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		jobs = append(jobs, Jobs{
			Title: e.Text,
			Link:  normalizeLink(link),
		})
	})

	// Job Timestamp
	c.OnHTML("tr > td.subtext > span.age", func(e *colly.HTMLElement) {
		timestamp := e.Attr("title")

		jobs[timestampIndex].Timestamp = timestamp
		timestampIndex++
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://news.ycombinator.com/jobs")

	jsonData, _ := json.MarshalIndent(jobs, "", " ")
	ioutil.WriteFile("jobs.json", jsonData, 0644)
}
