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

func LoginEmployee(employees models.Tab_Employee, t int) bool {
	userName, _ := pterm.DefaultInteractiveTextInput.Show("Please Input Your Username")
	password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please Input Your Password")

	for i := 0; i < t; i++ {
		if employees[i].UserName == userName && employees[i].Password == password {
			global.IsLogin = true

			return global.IsLogin
		}
	}

	return global.IsLogin
}

func RegisterEmployee(employees *models.Tab_Employee, t *int) bool {
	if *t < global.NMAX {
		fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Fullname")
		shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Shortname")
		userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Username")
		password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please input your Password")
		department, _ := pterm.DefaultInteractiveTextInput.Show("Please Input your Department")
		manager, _ := pterm.DefaultInteractiveTextInput.Show("Who is your Manager?")
		bio, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show("Tell Me About your Personal")

		employee := models.Employee_Model{
			Id:         *t,
			FullName:   fullName,
			ShortName:  shortName,
			UserName:   userName,
			Password:   password,
			Bio:        bio,
			Department: department,
			Manager:    manager,
			CreatedAt:  time.Now(),
		}

		employees[*t] = employee
		*t += 1

		global.IsRegistered = true
		return global.IsRegistered
	} else {
		pterm.Error.Println("You Have Reached The Maksimum Data Allowed!")
	}

	return global.IsRegistered
}

func CreateEmployee(employees *models.Tab_Employee, t *int) {
	if *t < global.NMAX {
		fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Fullname")
		shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Shortname")
		userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Username")
		password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please input your Password")
		department, _ := pterm.DefaultInteractiveTextInput.Show("Please Input your Department")
		manager, _ := pterm.DefaultInteractiveTextInput.Show("Who is your Manager?")
		bio, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show("Tell Me About your Personal")

		employee := models.Employee_Model{
			Id:         *t,
			FullName:   fullName,
			ShortName:  shortName,
			UserName:   userName,
			Password:   password,
			Bio:        bio,
			Department: department,
			Manager:    manager,
			CreatedAt:  time.Now(),
		}

		employees[*t] = employee
		*t += 1
	} else {
		pterm.Error.Println("You Have Reached The Maksimum Data Allowed!")
	}
}

func ShowAllEmployee(employees models.Tab_Employee, t int) {
	tableEmployee := table.New(os.Stdout)
	tableEmployee.SetHeaders("#", "Username", "Password", "Fullname", "Shortname", "Bio", "Department", "Manager")

	for i := 0; i < t; i++ {
		tableEmployee.AddRow(employees[i].UserName, employees[i].Password, employees[i].FullName,
			employees[i].ShortName, employees[i].Bio, employees[i].Department, employees[i].Manager)
	}

	tableEmployee.AddFooters("", "", "", "", "", "", "", "Total Data:"+strconv.Itoa(len(employees)))
	tableEmployee.Render()
}

func ShowAllEmployeeByDepartment(employees models.Tab_Employee, t int) {
	var departments []string

	tableEmployee := table.New(os.Stdout)
	tableEmployee.SetHeaders("#", "Department", "Manager", "Fullname", "Shortname", "Bio")

	for i := 0; i < t; i++ {
		departments = append(departments, employees[i].Department)
	}

	selectedDepartments, _ := pterm.DefaultInteractiveSelect.WithOptions(departments).Show(
		"Please Select Department to Filter")

	for j := 0; j < t; j++ {
		if employees[j].Department == selectedDepartments {
			tableEmployee.AddRow(employees[j].Department, employees[j].Manager, employees[j].FullName,
				employees[j].ShortName, employees[j].Bio)
		}
	}

	tableEmployee.AddFooters("", "", "", "", "", "Total Data:"+strconv.Itoa(len(employees)))
	tableEmployee.Render()
}

func EditByIdEmployee() {}

func DeleteByIdEmployee() {}
