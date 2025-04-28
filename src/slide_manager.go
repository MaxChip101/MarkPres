package main

import (
	"log"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
)

func GetMarkdownContent(file_name string) {

	contentBytes, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
		return
	}

	md := goldmark.New()
	reader := text.NewReader(contentBytes)
	doc := md.Parser().Parse(reader)
	if doc == nil {
		log.Fatal("Failed to parse Markdown")
		return
	}

}
