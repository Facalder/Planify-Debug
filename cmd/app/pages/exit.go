package pages

import (
	global "github.com/Facalder/Planify"
	"github.com/Facalder/Planify/components"
	"github.com/pterm/pterm"
)

func Exit() {
	components.TerminalBigTitle("Thanks!")
	components.TerminalHeaderSection("Have a nice day, Be Careful!")

	for {
		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(1) - Reload Project")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		global.ChooseMenu = choose

		switch global.ChooseMenu {
		case "1":
			components.Loader(global.ReloadingAppDuration, "Reloading Project, Please Wait a Second...", func() {
				IntroScreen()
			})
		default:
			pterm.DefaultLogger.Error("Menu Option Is Not Valid!, Please Fill Command Correctly!")
		}
	}
}
