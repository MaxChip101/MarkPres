package main

import (
	"log"
	"os"
	"strings"

	"golang.org/x/term"
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

	slides, err := GetMarkdownSlides(file)

	// convert to raw mode for input
	old_state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), old_state)

}
