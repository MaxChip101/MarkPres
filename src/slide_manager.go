package main

import (
	"errors"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

type Slide struct {
	content []ast.Node
}

/*
this function builds a slide based on a markdown file
returns a slide that contains the ast content of the markdown file
also returns if the slide is the last slide as a boolean
*/
func MakeSlide(doc ast.Node) (Slide, bool) {
	return Slide, false
}

/*
this function creates an array of slides for a presentation
returns the slide array
also returns any errors
*/
func GetMarkdownSlides(file_name string) ([]Slide, error) {

	contentBytes, err := os.ReadFile(file_name)
	if err != nil {
		return nil, err
	}

	md := goldmark.New()
	reader := text.NewReader(contentBytes)
	doc := md.Parser().Parse(reader)
	if doc == nil {
		return nil, errors.New("Failed to parse Markdown")
	}
	return nil, nil
}
