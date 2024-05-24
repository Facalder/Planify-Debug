package pages

import (
	global "github.com/Facalder/Planify"
	"github.com/Facalder/Planify/components"
	"github.com/Facalder/Planify/controller"
	"github.com/Facalder/Planify/models"
	"github.com/pterm/pterm"
)

func Employee() {
	components.TerminalHeaderSection("Welcome to Manager, Create account or Log in to get into Employee Dashboard")

	for {
		pterm.DefaultSection.Println("Have an account?")
		pterm.DefaultBasicText.Println("(1) - Login")

		pterm.DefaultSection.Println("Don't Have Any Account?")
		pterm.DefaultBasicText.Println("(2) - Create Account")

		pterm.DefaultSection.Println("Others Available Commands?")
		pterm.DefaultBasicText.Println("(B) - Back")
		pterm.DefaultBasicText.Println("(E) - Exit Program")

		choose, _ := pterm.DefaultInteractiveTextInput.Show("Choose Command")
		global.ChooseMenu = choose

		switch global.ChooseMenu {
		case "1":
			if controller.LoginEmployee(models.Employees, models.NEmployee) {
				components.Loader(global.LoadingDuration, "Validating Your Data, Please Wait a Second...", func() {
					MenuEmployee()
				})
			} else {
				components.Loader(global.ValidatingDuration, "Validating Your Data, Please Wait a Second...", func() {
					pterm.Error.Println("Username and Password Doesn't Match, Please Try again!")
				})
			}
		case "2":
			if controller.RegisterEmployee(&models.Employees, &models.NEmployee) {
				components.Loader(global.RegisteringDuration, "Create and Save Your Data...", func() {
					MenuEmployee()
				})
			} else {
				pterm.Error.Println("Can't Register, Please Try Again!")
			}
		case "B":
			components.Loader(global.LoadingDuration, "Leaving Employee...", func() {
				InitialMenu()
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
