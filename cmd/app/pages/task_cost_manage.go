package pages

import (
	global "github.com/Facalder/Planify"
	"github.com/Facalder/Planify/components"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func ManageTaskCost() {
	components.TerminalHeaderSection(putils.CenterText("Welcome to Manage Task (Admin) Menu!, \n" +
		"Manage Your Employees to Work Harder, and Harder!"))

	for {
		pterm.DefaultSection.Println("Available Commands")
		pterm.DefaultBasicText.Println("(1) - Show All Task Cost")
		pterm.DefaultBasicText.Println("(2) - Show All Average Task Cost")
		pterm.DefaultBasicText.Println("(3) - Show All The Highest Task Cost")
		pterm.DefaultBasicText.Println("(4) - Show All The Cheapest Task Cost")
		pterm.DefaultBasicText.Println("(5) - Show All Task From The Highest to Cheaper")
		pterm.DefaultBasicText.Println("(6) - Show All Task From The Cheaper to Highest")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(B) - Back")
		pterm.DefaultBasicText.Println("(E) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		global.ChooseMenu = choose

		switch global.ChooseMenu {
		case "1":

		case "2":

		case "3":

		case "4":

		case "10":

		case "B":
			components.Loader(global.LoadingDuration, "Leaving Manage Task Menu....", func() {
				ManageTask(func() {})
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
