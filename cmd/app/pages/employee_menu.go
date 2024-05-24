package pages

import (
	global "github.com/Facalder/Planify"
	"github.com/Facalder/Planify/components"
	"github.com/pterm/pterm"
)

func MenuEmployee() {
	components.TerminalBigTitle("Welcome to Employee Dashboard, Choose Command to use Functionality")

	for {
		pterm.DefaultSection.Println("Choose Manage Function")
		pterm.DefaultBasicText.Println("(1) - Manage Task")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(B) - Back")
		pterm.DefaultBasicText.Println("(E) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		global.ChooseMenu = choose

		switch global.ChooseMenu {
		case "1":
			components.Loader(global.LoadingDuration, "Loading...", func() {
				ManageTask(func() {
					MenuEmployee()
				})
			})

		case "B":
			components.Loader(global.LoadingDuration, "Leaving Employee Menu....", func() {
				Employee()
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
