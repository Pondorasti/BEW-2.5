package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type Article struct {
	Title      string
	Paragraphs []string
}

func parseFile(filePath string) Article {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	article := Article{}
	for index, line := range strings.Split(string(fileContents), "\n") {
		if index == 0 {
			article.Title = line
		} else if line != "" {
			article.Paragraphs = append(article.Paragraphs, line)
		}
	}

	return article
}

func generateHtml(article Article) {
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err := t.Execute(os.Stdout, article)
	if err != nil {
		panic(err)
	}
}

func main() {
	article := parseFile("data/first-post.txt")
	generateHtml(article)
}
