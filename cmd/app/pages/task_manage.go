package pages

import (
	global "github.com/Facalder/Planify"
	"github.com/Facalder/Planify/components"
	"github.com/Facalder/Planify/controller"
	"github.com/Facalder/Planify/models"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func ManageTask(menu func()) {
	components.TerminalHeaderSection(putils.CenterText("Welcome to Manage Task Menu!, \n" +
		"Manage Your Employees to Work Harder, and Harder!"))

	for {
		pterm.DefaultSection.Println("Available Commands")
		pterm.DefaultBasicText.Println("(1) - Add New Data Task")
		pterm.DefaultBasicText.Println("(2) - Show All Task")
		pterm.DefaultBasicText.Println("(3) - Show All Task by Ascending")
		pterm.DefaultBasicText.Println("(4) - Show All Task by Descending")
		pterm.DefaultBasicText.Println("(5) - Show Only Task by Deadline")
		pterm.DefaultBasicText.Println("(6) - Show Only Task by Priority")
		pterm.DefaultBasicText.Println("(7) - Show Only Task by Category")
		pterm.DefaultBasicText.Println("(8) - Search Available Task by Name")
		pterm.DefaultBasicText.Println("(9) - Edit Data Task")
		pterm.DefaultBasicText.Println("(10) - Remove Data Task")

		pterm.DefaultSection.Println("Others Task Commands")
		pterm.DefaultBasicText.Println("(TS) - Manage Task Cost")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(B) - Back")
		pterm.DefaultBasicText.Println("(E) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		global.ChooseMenu = choose

		switch global.ChooseMenu {
		case "1":
			controller.CreateTask(&models.Tasks, &models.NTask)
		case "2":
			controller.ShowAllTask(models.Tasks, models.NTask)
		case "3":
			controller.ShowAllTaskByASC(models.Tasks, models.NTask)
		case "4":
			controller.ShowAllTaskByDSC(models.Tasks, models.NTask)
		case "5":

		case "6":
			controller.ShowAllTaskByPriority(models.Tasks, models.NTask)
		case "7":
			controller.ShowAllTaskByCategory(models.Tasks, models.NTask)
		case "8":
			controller.SearchTaskByName(&models.Tasks, &models.NTask)
		case "9":
			controller.EditByIdTask(&models.Tasks, models.NTask)
		case "10":
			controller.DeleteByIdTask(&models.Tasks, &models.NTask)
		case "TS":
			components.Loader(global.LoadingDuration, "Loading...", func() {
				ManageTaskCost()
			})
		case "B":
			components.Loader(global.LoadingDuration, "Leaving Manage Task Menu....", func() {
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
