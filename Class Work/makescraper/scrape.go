package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Jobs struct {
	Title     string
	Link      string
	timestamp string
}

func normalizeLink(link string) string {
	if link[0] == 'i' { // item?id=
		link = "https://news.ycombinator.com/" + link
	}

	return link
}

func main() {
	data := []Jobs{}
	timestampIndex := 0
	c := colly.NewCollector()

	// Job Details
	c.OnHTML("tr.athing > td.title > a.titlelink", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		data = append(data, Jobs{
			Title: e.Text,
			Link:  normalizeLink(link),
		})
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	// Job Timestamp
	c.OnHTML("tr > td.subtext > span.age", func(e *colly.HTMLElement) {
		timestamp := e.Attr("title")
		fmt.Println(timestamp)

		data[timestampIndex].timestamp = timestamp
		timestampIndex++
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://news.ycombinator.com/jobs")

	fmt.Println("\n")
	fmt.Println(data[0])
}
