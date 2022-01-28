package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type Article struct {
	Title      string
	Paragraphs []string
}

func main() {
	fileContents, err := ioutil.ReadFile("data/first-post.txt")
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

	templateContents, err := ioutil.ReadFile("template.tmpl")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("template.tmpl").Parse(string(templateContents)))
	err = t.Execute(os.Stdout, article)
	if err != nil {
		fmt.Println("hello")
		panic(err)
	}
}
