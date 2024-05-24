package pages

import (
	global "github.com/Facalder/Planify"
	"github.com/Facalder/Planify/components"
	"github.com/Facalder/Planify/controller"
	"github.com/Facalder/Planify/models"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func ManageEmployee(menu func()) {
	components.TerminalHeaderSection(
		putils.CenterText("Welcome to Manage Employee Menu!, \n" +
			"Manage Your Employees to Work Harder, and Harder!"))

	for {
		pterm.DefaultSection.Println("Available Commands")
		pterm.DefaultBasicText.Println("(1) - Add New Data Employee")
		pterm.DefaultBasicText.Println("(2) - Edit Data Employee")
		pterm.DefaultBasicText.Println("(3) - Show All Available Employee")
		pterm.DefaultBasicText.Println("(4) - Show All Available Employee by Department")
		pterm.DefaultBasicText.Println("(5) - Search Available Employee by Name")
		pterm.DefaultBasicText.Println("(6) - Remove Data Employee")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(B) - Back")
		pterm.DefaultBasicText.Println("(E) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		global.ChooseMenu = choose

		switch global.ChooseMenu {
		case "1":
			controller.CreateEmployee(&models.Employees, &models.NEmployee)
		case "2":
			controller.EditByIdEmployee(&models.Employees, models.NEmployee)
		case "3":
			controller.ShowAllEmployee(models.Employees, models.NEmployee)
		case "4":
			controller.ShowAllEmployeeByDepartment(models.Employees, models.NEmployee)
		case "5":
			controller.SearchEmployeeByUsername(&models.Employees, &models.NEmployee)
		case "6":
			controller.DeleteByIdEmployee(&models.Employees, &models.NEmployee)
		case "B":
			components.Loader(global.LoadingDuration, "Leaving Manage Employee Menu....", func() {
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
