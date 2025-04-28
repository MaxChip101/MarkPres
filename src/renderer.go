package main

import (
	"fmt"

	"golang.org/x/term"
	"rsc.io/quote"
)

func GetTerminalSize() (int, int, error) {
	width, height, err := term.GetSize(0)
	return width, height, err
}

func Render(slide Slide) {
	fmt.Println(quote.Go())
}
