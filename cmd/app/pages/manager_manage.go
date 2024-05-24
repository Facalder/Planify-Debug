package pages

import (
	global "github.com/Facalder/Planify"
	"github.com/Facalder/Planify/components"
	"github.com/Facalder/Planify/controller"
	"github.com/Facalder/Planify/models"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func ManageManager(menu func()) {
	components.TerminalHeaderSection(putils.CenterText("Welcome to Manage Manager Menu!, \n" +
		"Manage Your Manager to Help Employees Complete Their Tasks Well"))

	for {
		pterm.DefaultSection.Println("Available Commands")
		pterm.DefaultBasicText.Println("(1) - Add New Data Manager")
		pterm.DefaultBasicText.Println("(2) - Edit Data Manager")
		pterm.DefaultBasicText.Println("(3) - Show All Available Manager")
		pterm.DefaultBasicText.Println("(4) - Show All Available Manager by Department")
		pterm.DefaultBasicText.Println("(5) - Remove Data Manager")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(B) - Back")
		pterm.DefaultBasicText.Println("(E) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		global.ChooseMenu = choose

		switch global.ChooseMenu {
		case "1":
			controller.CreateManager(&models.Managers, &models.NManager)
		case "2":

		case "3":
			controller.ShowAllManager(models.Managers, models.NManager)
		case "4":
			controller.ShowAllManagerByDepartment(models.Managers, models.NManager)
		case "5":

		case "B":
			components.Loader(global.LoadingDuration, "Leaving Manage Manager Menu....", func() {
				menu()
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
