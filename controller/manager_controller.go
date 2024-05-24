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

		components.Loader(global.SaveDataDuration, "Create and Save Your Data...", func() {
			managers[*t] = manager
			*t += 1

			pterm.Info.Println("Manager Data Has Been Created!")
		})
	} else {
		pterm.Error.Println("You Have Reached The Maksimum Data Allowed!")
	}
}

func ShowAllManager(managers models.Tab_Manager, t int) {
	if t == 0 {
		pterm.Error.Println("There's No Manager Data, Please Create The New One!")
	} else {
		tableManager := table.New(os.Stdout)
		tableManager.SetHeaders("#", "Username", "Password", "Fullname", "Shortname", "Bio", "Department")

		for i := 0; i < t; i++ {
			tableManager.AddRow(strconv.Itoa(managers[i].Id), managers[i].UserName, managers[i].Password, managers[i].FullName,
				managers[i].ShortName, managers[i].Bio, managers[i].Department)
		}

		components.Loader(global.GetAllDataDuration, "Consuming Data, Please Wait...", func() {
			tableManager.Render()
		})
	}
}

func ShowAllManagerByDepartment(managers models.Tab_Manager, t int) {
	if t == 0 {
		pterm.Error.Println("There's No Manager Data, Please Create The New One!")
	} else {
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
				tableManager.AddRow(strconv.Itoa(managers[j].Id), managers[j].Department, managers[j].FullName,
					managers[j].ShortName, managers[j].Bio)
			}
		}

		components.Loader(global.SaveDataDuration, "Consuming Filtered Data, Please Wait...", func() {
			tableManager.Render()
		})
	}
}

func SearchManagerByUsername(managers *models.Tab_Manager, t *int) {
	var isFound bool = false

	tableManager := table.New(os.Stdout)
	tableManager.SetHeaders("#", "Username", "Password", "Fullname", "Shortname", "Bio", "Department")

	pterm.DefaultSection.Println("Search Manager by Their Username")
	userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Username")

	for i := 0; i < *t; i++ {
		if lib.CompareIgnoreCase(managers[i].UserName, userName) {
			isFound = true
			tableManager.AddRow(strconv.Itoa(managers[i].Id), managers[i].UserName, managers[i].Password, managers[i].FullName,
				managers[i].ShortName, managers[i].Bio, managers[i].Department)
		}
	}

	components.Loader(global.SearchDataDuration, "Search Manager, Please Wait...", func() {
		tableManager.Render()
	})

	if !isFound {
		pterm.Error.Println("Manager Name Not Found, Please Input Another Name!")
	}
}

func EditByIdManager(managers *models.Tab_Manager, t int) {
	if t == 0 {
		pterm.Error.Println("There's No Manager Data, Please Create The New One!")
	} else {
		var Id int

		tableManager := table.New(os.Stdout)
		tableManager.SetHeaders("#", "Username", "Password", "Fullname", "Shortname", "Bio", "Department")

		for i := 0; i < t; i++ {
			tableManager.AddRow(strconv.Itoa(managers[i].Id), managers[i].UserName, managers[i].Password, managers[i].FullName,
				managers[i].ShortName, managers[i].Bio, managers[i].Department)
		}

		pterm.DefaultSection.Println("Choose Id Where you Want to Edit")
		idStr, _ := pterm.DefaultInteractiveTextInput.Show("Please Input Manager Id")
		id, _ := strconv.Atoi(idStr)

		if Id != -1 && Id >= 0 && Id == id {
			pterm.DefaultSection.Println("Available Data to Edit")
			fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your New Fullname")
			shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your New Shortname")
			bio, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show("Tell Me Again About your Personal")

			components.Loader(global.SaveOnAfterEditDuration, "Editing and Save The Data, Please Wait...", func() {
				managers[Id].FullName = fullName
				managers[Id].ShortName = shortName
				managers[Id].Bio = bio
			})
		}
	}
}

func DeleteByIdManager(managers *models.Tab_Manager, t *int) {
	if *t == 0 {
		pterm.Error.Println("There's No Manager Data, Please Create The New One!")
	} else {
		var Id int

		tableManager := table.New(os.Stdout)
		tableManager.SetHeaders("#", "Username", "Password", "Fullname", "Shortname", "Bio", "Department")

		for i := 0; i < *t; i++ {
			tableManager.AddRow(strconv.Itoa(managers[i].Id), managers[i].UserName, managers[i].Password, managers[i].FullName,
				managers[i].ShortName, managers[i].Bio, managers[i].Department)

			Id = managers[i].Id
		}

		pterm.DefaultSection.Println("Choose Id Where you Want to Delete")
		idStr, _ := pterm.DefaultInteractiveTextInput.Show("Please Input Manager Id")
		id, _ := strconv.Atoi(idStr)

		if Id != -1 && id >= 0 && Id == id {
			components.Loader(global.SaveOnAfterDeleteDuration, "Deleting The Data, Please Wait...", func() {
				managers[*t-1] = models.Manager_Model{}
				*t = -1
			})

			pterm.Info.Println("Data Manager Has Been Deleted!")
		} else {
			pterm.Error.Println("There's No Manager Data")
		}
	}
}
