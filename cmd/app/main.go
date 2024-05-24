package main

import (
	"github.com/Facalder/Planify/cmd/app/pages"
)

func main() {
	pages.IntroScreen()
	print("\033[H\033[2J")
}
