package controller

import (
	global "github.com/Facalder/Planify"
	"github.com/Facalder/Planify/models"
	"github.com/aquasecurity/table"
	"github.com/pterm/pterm"
	"os"
	"strconv"
	"time"
)

func LoginManager(managers models.Tab_Manager, t int) bool {
	userName, _ := pterm.DefaultInteractiveTextInput.Show("Please Input Your Username")
	password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please Input Your Password")

	for i := 0; i < t; i++ {
		if managers[i].UserName == userName && managers[i].Password == password {
			global.IsLogin = true

			return global.IsLogin
		}
	}

	return global.IsLogin
}

func RegisterManager(managers *models.Tab_Manager, t *int) bool {
	if *t < global.NMAX {
		fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Fullname")
		shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Shortname")
		userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Username")
		password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please input your Password")
		department, _ := pterm.DefaultInteractiveTextInput.Show("Please Input your Department")
		bio, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show("Tell Me About your Personal")

		manager := models.Manager_Model{
			Id:         *t,
			FullName:   fullName,
			ShortName:  shortName,
			UserName:   userName,
			Password:   password,
			Bio:        bio,
			Department: department,
			CreatedAt:  time.Now(),
		}

		managers[*t] = manager
		*t += 1

		global.IsRegistered = true
		return global.IsRegistered
	} else {
		pterm.Error.Println("You Have Reached The Maksimum Data Allowed!")
	}

	return global.IsRegistered
}

func CreateManager(managers *models.Tab_Manager, t *int) {
	if *t < global.NMAX {
		fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Fullname")
		shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Shortname")
		userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Username")
		password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please input your Password")
		department, _ := pterm.DefaultInteractiveTextInput.Show("Please Input your Department")
		bio, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show("Tell Me About your Personal")

		manager := models.Manager_Model{
			Id:         *t,
			FullName:   fullName,
			ShortName:  shortName,
			UserName:   userName,
			Password:   password,
			Bio:        bio,
			Department: department,
			CreatedAt:  time.Now(),
		}

		managers[*t] = manager
		*t += 1
	} else {
		pterm.Error.Println("You Have Reached The Maksimum Data Allowed!")
	}
}

func ShowAllManager(managers models.Tab_Manager, t int) {
	tableManager := table.New(os.Stdout)
	tableManager.SetHeaders("#", "Username", "Password", "Fullname", "Shortname", "Bio", "Department")

	for i := 0; i < t; i++ {
		tableManager.AddRow(managers[i].UserName, managers[i].Password, managers[i].FullName,
			managers[i].ShortName, managers[i].Bio, managers[i].Department)
	}

	tableManager.AddFooters("", "", "", "", "", "", "", "Total Data:"+strconv.Itoa(len(managers)))
	tableManager.Render()
}

func ShowAllManagerByDepartment(managers models.Tab_Manager, t int) {
	var departments []string

	tableManager := table.New(os.Stdout)
	tableManager.SetHeaders("#", "Department", "Fullname", "Shortname", "Bio")

	for i := 0; i < t; i++ {
		departments = append(departments, managers[i].Department)
	}

	selectedDepartments, _ := pterm.DefaultInteractiveSelect.WithOptions(departments).Show(
		"Please Select Department to Filter")

	for j := 0; j < t; j++ {
		if managers[j].Department == selectedDepartments {
			tableManager.AddRow(managers[j].Department, managers[j].FullName,
				managers[j].ShortName, managers[j].Bio)
		}
	}

	tableManager.AddFooters("", "", "", "", "", "Total Data:"+strconv.Itoa(len(managers)))
	tableManager.Render()
}

func EditByIdManager() {}

func DeleteByIdManager() {}
