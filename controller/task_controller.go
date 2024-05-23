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

		tasks[*t] = task
		*t += 1
	} else {
		pterm.Error.Println("You Have Reached The Maximum Number of Tasks!")
	}
}

func ShowAllTask(tasks models.Tab_Task, t int) {
	tableTask := table.New(os.Stdout)
	tableTask.SetHeaders("#", "Name", "Description", "Priority", "Status", "Notes", "Category")

	for i := 0; i < t; i++ {
		tableTask.AddRow(strconv.Itoa(tasks[i].Id), tasks[i].Name, tasks[i].Description,
			tasks[i].Priority, tasks[i].Status, tasks[i].Notes, tasks[i].Category)
	}

	tableTask.AddFooters("", "", "", "", "", "", "Total Data:"+strconv.Itoa(len(tasks)))
	tableTask.Render()
}

func ShowAllTaskByASC() {}

func ShowAllTaskByDSC() {}

func ShowAllTaskByDeadline() {}

func ShowAllTaskByPriority() {}

func ShowAllTaskByCategory() {}
