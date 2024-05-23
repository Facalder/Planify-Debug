package page

import (
	global "github.com/Facalder/Planify"
	"github.com/Facalder/Planify/components"
	"github.com/pterm/pterm"
)

func MenuAdmin() {
	components.TerminalBigTitle("Welcome to Admin Dashboard, Choose Command to use Functionality")

	for {
		pterm.DefaultSection.Println("Choose Manage Function")
		pterm.DefaultBasicText.Println("(1) - Manage Manager")
		pterm.DefaultBasicText.Println("(2) - Manage Employee")
		pterm.DefaultBasicText.Println("(3) - Manage Task")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(B) - Back")
		pterm.DefaultBasicText.Println("(E) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		global.ChooseMenu = choose

		switch global.ChooseMenu {
		case "1":
			components.Loader(global.LoadingDuration, "Loading...", func() {

			})

		case "2":
			components.Loader(global.LoadingDuration, "Loading...", func() {

			})

		case "3":
			components.Loader(global.LoadingDuration, "Loading...", func() {

			})

		case "B":
			components.Loader(global.LoadingDuration, "Leaving Admin Menu....", func() {
				Admin()
			})

		case "E":
			components.Loader(global.ExitingProgramDuration, "Exiting Program, Please Wait a Second...", func() {
				Exit()
			})

		default:
			pterm.Error.Println("Menu Option Is Not Valid!, Please Fill Command Correctly!")
		}
	}
}
