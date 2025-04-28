package main

import (
	"golang.org/x/term"
)

func GetTerminalSize() (int, int, error) {
	width, height, err := term.GetSize(0)
	return width, height, err
}

/*
this function takes a slide and renders on the terminal depending on the markdown ast in the slide
returns nothing as it doesn't need to return anything
*/
func Render(slide Slide) {

}
