package controller

import (
	global "github.com/Facalder/Planify"
	"github.com/Facalder/Planify/components"
	"github.com/Facalder/Planify/lib"
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

		components.Loader(global.SaveDataDuration, "Create and Save Your Data...", func() {
			employees[*t] = employee
			*t += 1

			pterm.Info.Println("Employee Data Has Been Created!")
		})
	} else {
		pterm.Error.Println("You Have Reached The Maksimum Data Allowed!")
	}
}

func ShowAllEmployee(employees models.Tab_Employee, t int) {
	if t == 0 && t < 0 {
		pterm.Error.Println("There's No Employee Data, Please Create The New One!")
	} else {
		tableEmployee := table.New(os.Stdout)
		tableEmployee.SetHeaders("#", "Username", "Password", "Department", "Manager", "Fullname", "Shortname", "Bio")

		for i := 0; i < t; i++ {
			tableEmployee.AddRow(strconv.Itoa(employees[i].Id), employees[i].UserName, employees[i].Password,
				employees[i].Department, employees[i].Manager, employees[i].FullName, employees[i].ShortName, employees[i].Bio)
		}

		components.Loader(global.GetAllDataDuration, "Consuming Data, Please Wait...", func() {
			tableEmployee.Render()
		})
	}
}

func ShowAllEmployeeByDepartment(employees models.Tab_Employee, t int) {
	if t == 0 {
		pterm.Error.Println("There's No Employee Data, Please Create The New One!")
	} else {
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
				tableEmployee.AddRow(strconv.Itoa(employees[j].Id), employees[j].Department, employees[j].Manager, employees[j].FullName,
					employees[j].ShortName, employees[j].Bio)
			}
		}

		components.Loader(global.SaveDataDuration, "Consuming Filtered Data, Please Wait...", func() {
			tableEmployee.Render()
		})
	}
}

func SearchEmployeeByUsername(employees *models.Tab_Employee, t *int) {
	var isFound bool = false

	tableEmployee := table.New(os.Stdout)
	tableEmployee.SetHeaders("#", "Username", "Password", "Department", "Manager", "Fullname", "Shortname", "Bio")

	pterm.DefaultSection.Println("Search Employee by Their Username")
	userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Username")

	for i := 0; i < *t; i++ {
		if lib.CompareIgnoreCase(employees[i].UserName, userName) {
			isFound = true
			tableEmployee.AddRow(strconv.Itoa(employees[i].Id), employees[i].UserName, employees[i].Password,
				employees[i].Department, employees[i].Manager, employees[i].FullName, employees[i].ShortName, employees[i].Bio)
		}
	}

	components.Loader(global.SearchDataDuration, "Search Employee, Please Wait...", func() {
		tableEmployee.Render()
	})

	if !isFound {
		pterm.Error.Println("Employee Name Not Found, Please Input Another Name!")
	}
}

func EditByIdEmployee(employees *models.Tab_Employee, t int) {
	if t == 0 {
		pterm.Error.Println("There's No Employee Data, Please Create The New One!")
	} else {
		var Id int

		tableEmployee := table.New(os.Stdout)
		tableEmployee.SetHeaders("#", "Username", "Password", "Department", "Manager", "Fullname", "Shortname", "Bio")

		for i := 0; i < t; i++ {
			tableEmployee.AddRow(strconv.Itoa(employees[i].Id), employees[i].UserName, employees[i].Password,
				employees[i].Department, employees[i].Manager, employees[i].FullName, employees[i].ShortName, employees[i].Bio)

			Id = employees[i].Id
		}

		tableEmployee.Render()

		pterm.DefaultSection.Println("Choose Id Where you Want to Edit")
		idStr, _ := pterm.DefaultInteractiveTextInput.Show("Choose Id Wants to Edit")
		id, _ := strconv.Atoi(idStr)

		if Id != -1 && Id >= 0 && Id == id {
			pterm.DefaultSection.Println("Available Data to Edit")
			fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your New Fullname")
			shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your New Shortname")
			bio, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show("Tell Me Again About your Personal")

			components.Loader(global.SaveOnAfterEditDuration, "Editing and Save The Data, Please Wait...", func() {
				employees[Id].FullName = fullName
				employees[Id].ShortName = shortName
				employees[Id].Bio = bio
			})

			pterm.Info.Printfln("Data With %d Has Been Edited Successfully!", Id)
		} else {
			pterm.Error.Println("Id Not Found, Please Input Another Id!")
		}
	}
}

func DeleteByIdEmployee(employees *models.Tab_Employee, t *int) {
	if *t == 0 {
		pterm.Error.Println("There's No Employee Data, Please Create The New One!")
	} else {
		var Id int

		tableEmployee := table.New(os.Stdout)
		tableEmployee.SetHeaders("#", "Username", "Password", "Department", "Manager", "Fullname", "Shortname", "Bio")

		for i := 0; i < *t; i++ {
			tableEmployee.AddRow(strconv.Itoa(employees[i].Id), employees[i].UserName, employees[i].Password,
				employees[i].Department, employees[i].Manager, employees[i].FullName, employees[i].ShortName, employees[i].Bio)

			Id = employees[i].Id
		}

		tableEmployee.Render()

		pterm.DefaultSection.Println("Choose Id Where you Want to Delete")
		idStr, _ := pterm.DefaultInteractiveTextInput.Show("Choose Id Wants to Delete")
		id, _ := strconv.Atoi(idStr)

		if Id != -1 && id >= 0 && Id == id {
			components.Loader(global.SaveOnAfterDeleteDuration, "Deleting The Data, Please Wait...", func() {
				employees[*t-1] = models.Employee_Model{}
				*t = -1
			})

			pterm.Info.Println("Data Employe Has Been Deleted!")
		} else {
			pterm.Error.Println("There's No Employee Data")
		}
	}
}
