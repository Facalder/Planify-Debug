package components

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func TerminalBigTitle(title string) {
	planifyLogo, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle(title, pterm.FgLightMagenta.ToStyle()),
	).Srender()

	pterm.DefaultCenter.Print(planifyLogo)
}

func TerminalHeaderSection(title string) {
	newHeader := pterm.HeaderPrinter{
		TextStyle:       pterm.NewStyle(pterm.FgLightWhite, pterm.Bold),
		BackgroundStyle: pterm.NewStyle(pterm.BgLightBlue),
		Margin:          1,
	}

	newHeader.WithFullWidth().Println(title)
}
