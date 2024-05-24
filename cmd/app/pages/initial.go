package pages

import (
	global "github.com/Facalder/Planify"
	"github.com/Facalder/Planify/components"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func InitialMenu() {
	components.TerminalHeaderSection(putils.CenterText("The beginning of the story of this program begins here!," +
		"\nChoose Your Role Before Use Application"))

	for {
		pterm.DefaultSection.Println("Pick Your Role?")
		pterm.DefaultBasicText.Println("(1) - Admin")
		pterm.DefaultBasicText.Println("(2) - Manager")
		pterm.DefaultBasicText.Println("(3) - Employee")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(E) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		global.ChooseMenu = choose

		switch global.ChooseMenu {
		case "1":
			components.Loader(global.LoadingDuration, "Entering Admin...", func() {
				Admin()
			})
		case "2":
			components.Loader(global.LoadingDuration, "Entering Manager...", func() {
				Manager()
			})

		case "3":
			components.Loader(global.LoadingDuration, "Entering Employee...", func() {
				Employee()
			})

		case "E":
			components.Loader(global.LoadingDuration, "Exiting Program, Please Wait a Second...", func() {
				Exit()
			})

		default:
			pterm.Error.Println("Menu Option Is Not Valid!, Please Fill Command Correctly!")
		}
	}
}
