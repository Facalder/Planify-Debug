package pages

import (
	global "github.com/Facalder/Planify"
	"github.com/Facalder/Planify/components"
	"github.com/pterm/pterm"
	"time"
)

func IntroScreen() {
	components.TerminalBigTitle("planify")

	components.TerminalHeaderSection("Welcome to Planify!, the Task Management App - Made with love for you!")

	pterm.Info.Println("This project was generated with the latest version of PTerm and Golang!" +
		"\nPlanify works on nearly every terminal and operating system." +
		"\nIt's super easy to use!" +
		"\nIf you want the sourcecode, hmm this project is not open source 😭" +
		"\nYou can see the code of this demo in the " + pterm.LightMagenta("github.com/Facalder/Planify") +
		"\nThis page was updated at: " + pterm.Green(time.Now().Format("02 Jan 2006 - 15:04:05 MST")))

	components.Loader(global.InitAppDuration, "Waiting init for 3 seconds", func() {
		InitialMenu()
	})
}
