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

func CreateTask(tasks *models.Tab_Task, t *int) {
	if *t < global.NMAX {
		name, _ := pterm.DefaultInteractiveTextInput.Show("Please Input your Task Name")
		description, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show("Please Input your Task Description")
		priority, _ := pterm.DefaultInteractiveTextInput.Show("Please Input your Task Priority, answer[low, medium, high]")
		status := "not-completed"
		cost, _ := pterm.DefaultInteractiveTextInput.Show("Please Input Task Cost")
		notes, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show("Task Note")
		category, _ := pterm.DefaultInteractiveTextInput.Show("Please Input your Task Category")

		parseCost, _ := strconv.ParseFloat(cost, 64)

		switch priority {
		case "medium":
			pterm.Yellow(priority)
		case "high":
			pterm.Red(priority)
		default:
			pterm.Gray(priority)
		}

		task := models.Task_Model{
			Id:          *t,
			Name:        name,
			Description: description,
			Priority:    priority,
			Status:      status,
			Cost:        parseCost,
			Notes:       notes,
			Category:    category,
			CreatedAt:   time.Now(),
		}

		components.Loader(global.SaveDataDuration, "Create and Save your Data...", func() {
			tasks[*t] = task
			*t += 1
		})
	} else {
		pterm.Error.Println("You Have Reached The Maximum Number of Tasks!")
	}
}

func ShowAllTask(tasks models.Tab_Task, t int) {
	if t == 0 {
		pterm.Error.Println("There's No Task Data, Please Create The New One!")
	} else {
		tableTask := table.New(os.Stdout)
		tableTask.SetHeaders("#", "Name", "Description", "Priority", "Status", "Notes", "Category", "Cost")

		for i := 0; i < t; i++ {
			tableTask.AddRow(strconv.Itoa(tasks[i].Id), tasks[i].Name, tasks[i].Description,
				tasks[i].Priority, tasks[i].Status, tasks[i].Notes, tasks[i].Category, strconv.FormatFloat(tasks[i].Cost, 'f', 2, 64))
		}

		components.Loader(global.GetAllDataDuration, "Consuming Data, Please Wait...", func() {
			tableTask.Render()
		})
	}
}

func ShowAllTaskByASC(tasks models.Tab_Task, t int) {
	if t == 0 {
		pterm.Error.Println("There's No Task Data, Please Create The New One!")
	} else {
		for i := 0; i < t-1; i++ {
			for j := 0; j < t-i-1; j++ {
				if !lib.CompareIgnoreCase(tasks[j].Name, tasks[j+1].Name) {
					temp := tasks[j]
					tasks[j] = tasks[j+1]
					tasks[j+1] = temp
				}
			}
		}

		tableTask := table.New(os.Stdout)
		tableTask.SetHeaders("#", "Name (ASC)", "Description", "Priority", "Status", "Notes", "Category", "Cost")

		for n := 0; n < t; n++ {
			tableTask.AddRow(strconv.Itoa(tasks[n].Id), tasks[n].Name, tasks[n].Description,
				tasks[n].Priority, tasks[n].Status, tasks[n].Notes, tasks[n].Category, strconv.FormatFloat(tasks[n].Cost, 'f', 2, 64))
		}

		components.Loader(global.GetAllDataDuration, "Sorting Data by Ascending, Please Wait...", func() {
			tableTask.Render()
		})
	}
}

func ShowAllTaskByDSC(tasks models.Tab_Task, t int) {
	if t == 0 {
		pterm.Error.Println("There's No Task Data, Please Create The New One!")
	} else {
		for i := 0; i < t-1; i++ {
			for j := 0; j < t-i-1; j++ {
				if lib.CompareIgnoreCase(tasks[j].Name, tasks[j+1].Name) {
					temp := tasks[j]
					tasks[j] = tasks[j+1]
					tasks[j+1] = temp
				}
			}
		}

		tableTask := table.New(os.Stdout)
		tableTask.SetHeaders("#", "Name (DSC)", "Description", "Priority", "Status", "Notes", "Category", "Cost")

		for n := 0; n < t; n++ {
			tableTask.AddRow(strconv.Itoa(tasks[n].Id), tasks[n].Name, tasks[n].Description,
				tasks[n].Priority, tasks[n].Status, tasks[n].Notes, tasks[n].Category, strconv.FormatFloat(tasks[n].Cost, 'f', 2, 64))
		}

		components.Loader(global.GetAllDataDuration, "Sorting Data by Ascending, Please Wait...", func() {
			tableTask.Render()
		})
	}
}

func ShowAllTaskByDeadline() {}

func ShowAllTaskByPriority(tasks models.Tab_Task, t int) {
	if t == 0 {
		pterm.Error.Println("There's No Task Data, Please Create The New One!")
	} else {
		var priorities []string

		for i := 0; i < t; i++ {
			priorities = append(priorities, tasks[i].Priority)
		}

		selectedPrority, _ := pterm.DefaultInteractiveSelect.WithOptions(priorities).Show("Show Only Priority Selected")

		tableTask := table.New(os.Stdout)
		tableTask.SetHeaders("#", "Priority", "Name", "Description", "Status", "Notes", "Category", "Cost")

		for j := 0; j < t; j++ {
			if tasks[j].Priority == selectedPrority {
				tableTask.AddRow(strconv.Itoa(tasks[j].Id), tasks[j].Priority, tasks[j].Name,
					tasks[j].Description, tasks[j].Status, tasks[j].Notes, tasks[j].Category, strconv.FormatFloat(tasks[j].Cost, 'f', 2, 64))
			}
		}

		components.Loader(global.SearchDataDuration, "Search Task, Please Wait...", func() {
			tableTask.Render()
		})
	}
}

func ShowAllTaskByCategory(tasks models.Tab_Task, t int) {
	if t == 0 {
		pterm.Error.Println("There's No Task Data, Please Create The New One!")
	} else {
		var categories []string

		for i := 0; i < t; i++ {
			categories = append(categories, tasks[i].Category)
		}

		selectedCategory, _ := pterm.DefaultInteractiveSelect.WithOptions(categories).Show("Show Only Priority Selected")

		tableTask := table.New(os.Stdout)
		tableTask.SetHeaders("#", "Category", "Name", "Description", "Priority", "Status", "Notes", "Category", "Cost")

		for j := 0; j < t; j++ {
			if tasks[j].Priority == selectedCategory {
				tableTask.AddRow(strconv.Itoa(tasks[j].Id), tasks[j].Category, tasks[j].Name,
					tasks[j].Description, tasks[j].Priority, tasks[j].Status, tasks[j].Notes, tasks[j].Category, strconv.FormatFloat(tasks[j].Cost, 'f', 2, 64))
			}
		}

		components.Loader(global.SearchDataDuration, "Search Task, Please Wait...", func() {
			tableTask.Render()
		})
	}
}

func CompletedTask() {}

func SearchTaskByName(tasks *models.Tab_Task, t *int) {
	var isFound bool = false

	tableTask := table.New(os.Stdout)
	tableTask.SetHeaders("#", "Name", "Description", "Priority", "Status", "Notes", "Category", "Cost")

	pterm.DefaultSection.Println("Search Task by Their Name")
	name, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Task Name")

	for i := 0; i < *t; i++ {
		if lib.CompareIgnoreCase(tasks[i].Name, name) {
			isFound = true
			tableTask.AddRow(strconv.Itoa(tasks[i].Id), tasks[i].Name, tasks[i].Description,
				tasks[i].Priority, tasks[i].Status, tasks[i].Notes, tasks[i].Category, strconv.FormatFloat(tasks[i].Cost, 'f', 2, 64))
		}
	}

	components.Loader(global.SearchDataDuration, "Search Task, Please Wait...", func() {
		tableTask.Render()
	})

	if !isFound {
		pterm.Error.Println("Task Name Not Found, Please Input Another Name!")
	}
}

func EditByIdTask(tasks *models.Tab_Task, t int) {
	if t == 0 {
		pterm.Error.Println("There's No Task Data, Please Create The New One!")
	} else {
		var Id int

		tableTask := table.New(os.Stdout)
		tableTask.SetHeaders("#", "Name", "Description", "Priority", "Status", "Notes", "Category", "Cost")

		for i := 0; i < t; i++ {
			tableTask.AddRow(strconv.Itoa(tasks[i].Id), tasks[i].Name, tasks[i].Description,
				tasks[i].Priority, tasks[i].Status, tasks[i].Notes, tasks[i].Category, strconv.FormatFloat(tasks[i].Cost, 'f', 2, 64))

			Id = tasks[i].Id
		}

		pterm.DefaultSection.Println("Choose Id Where you Want to Edit")
		idStr, _ := pterm.DefaultInteractiveTextInput.Show("Please Input Task Id")
		id, _ := strconv.Atoi(idStr)

		if Id != -1 && Id >= 0 && Id == id {
			name, _ := pterm.DefaultInteractiveTextInput.Show("Please Input your Task Name")
			description, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show("Please Input your Task Description")
			priority, _ := pterm.DefaultInteractiveTextInput.Show("Please Input your Task Priority, answer[low, medium, high]")
			status := "not-completed"
			cost, _ := pterm.DefaultInteractiveTextInput.Show("Please Input Task Cost")
			notes, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show("Task Note")
			category, _ := pterm.DefaultInteractiveTextInput.Show("Please Input your Task Category")

			components.Loader(global.SaveOnAfterEditDuration, "Editing and Save The Data, Please Wait...", func() {
				tasks[id].Name = name
				tasks[id].Description = description
				tasks[id].Priority = priority
				tasks[id].Status = status
				tasks[id].Cost, _ = strconv.ParseFloat(cost, 64)
				tasks[id].Notes = notes
				tasks[id].Category = category
			})

			pterm.Info.Printfln("Data With %d Has Been Edited Successfully!", Id)
		} else {
			pterm.Error.Println("Id Not Found, Please Input Another Id!")
		}
	}
}

func DeleteByIdTask(tasks *models.Tab_Task, t *int) {
	if *t == 0 {
		pterm.Error.Println("There's No Task Data, Please Create The New One!")
	} else {
		var Id int

		tableTask := table.New(os.Stdout)
		tableTask.SetHeaders("#", "Name", "Description", "Priority", "Status", "Notes", "Category", "Cost")

		for i := 0; i < *t; i++ {
			tableTask.AddRow(strconv.Itoa(tasks[i].Id), tasks[i].Name, tasks[i].Description,
				tasks[i].Priority, tasks[i].Status, tasks[i].Notes, tasks[i].Category, strconv.FormatFloat(tasks[i].Cost, 'f', 2, 64))

			Id = tasks[i].Id
		}

		pterm.DefaultSection.Println("Choose Id Where you Want to Delete")
		idStr, _ := pterm.DefaultInteractiveTextInput.Show("Please Input Task Id")
		id, _ := strconv.Atoi(idStr)

		if Id != 1 && id >= 0 && Id == id {
			components.Loader(global.SaveOnAfterDeleteDuration, "Deleting The Data, Please Wait...", func() {
				tasks[*t-1] = models.Task_Model{}
				*t = -1
			})

			pterm.Info.Println("Data Task Has Been Deleted!")
		} else {
			pterm.Error.Println("There's No Employee Data")
		}
	}
}
