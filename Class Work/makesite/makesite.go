package main

import (
	"flag"
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

func generateHtml(article Article, fileName string) {
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	output, err := os.Create("dist/" + fileName + ".html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(output, article)
	if err != nil {
		panic(err)
	}
}

func main() {
	filePath := flag.String("file", "data/first-post.txt", "The path to the file to parse")
	flag.Parse()

	splitPath := strings.Split(*filePath, "/")
	fileName := splitPath[len(splitPath)-1]
	splitExtension := strings.Split(fileName, ".")
	fileName = splitExtension[0]

	fmt.Println(*filePath)
	fmt.Println(fileName)

	article := parseFile(*filePath)
	generateHtml(article, fileName)
}
