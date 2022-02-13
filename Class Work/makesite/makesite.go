package main

import (
	"flag"
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
	filePath := flag.String("file", "", "Parse file from the given path.")
	dirPath := flag.String("dir", "", "Parse all files from the given directory.")
	flag.Parse()

	if *filePath != "" {
		splitPath := strings.Split(*filePath, "/")
		fileName := splitPath[len(splitPath)-1]
		fileName = strings.Split(fileName, ".")[0]

		article := parseFile(*filePath)
		generateHtml(article, fileName)
	}

	if *dirPath != "" {
		files, err := ioutil.ReadDir(*dirPath)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".txt") {
				fileName := strings.Split((file.Name()), ".")[0]

				article := parseFile(*dirPath + "/" + file.Name())
				generateHtml(article, fileName)
			}
		}
	}
}

// Example Usage
// go run makesite -file=data/first-post.txt
// go run makesite -dir=data
