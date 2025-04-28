package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	var file string
	for arg := range args {
		switch args[arg] {
		case "--presenter-view":

		default:
			file = args[arg]
		}
	}

	_, found := strings.CutSuffix(file, ".md")

	if file == "" || !found {
		log.Fatal("mpres: No markdown file specified")
		return
	}

	GetMarkdownContent(file)
}
