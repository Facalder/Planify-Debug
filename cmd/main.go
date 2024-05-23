package main

import "github.com/Facalder/Planify/cmd/page"

func main() {
	page.IntroScreen()
	print("\033[H\033[2J")
}
