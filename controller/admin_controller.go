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

func LoginAdmin(admins models.Tab_Admin, t int) bool {
	userName, _ := pterm.DefaultInteractiveTextInput.Show("Please Input Your Username")
	password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please Input Your Password")

	for i := 0; i < t; i++ {
		if admins[i].UserName == userName && admins[i].Password == password {
			global.IsLogin = true

			return global.IsLogin
		}
	}

	return global.IsLogin
}

func RegisterAdmin(admins *models.Tab_Admin, t *int) bool {
	if *t < global.NMAX {
		fullName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Fullname")
		shortName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Shortname")
		userName, _ := pterm.DefaultInteractiveTextInput.Show("Please input your Username")
		password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Please input your Password")

		admin := models.Admin_Model{
			Id:        *t,
			FullName:  fullName,
			ShortName: shortName,
			UserName:  userName,
			Password:  password,
			CreatedAt: time.Now(),
		}

		admins[*t] = admin
		*t += 1

		global.IsRegistered = true
		return global.IsRegistered
	} else {
		pterm.Error.Println("You Have Reached The Maksimum Data Allowed!")
	}

	return global.IsRegistered
}

func ShowAllAdmin(admins models.Tab_Admin, t int) {
	tableAdmin := table.New(os.Stdout)
	tableAdmin.SetHeaders("#", "Username", "Password", "Fullname", "Shortname")

	for i := 0; i < t; i++ {
		tableAdmin.AddRow(admins[i].UserName, admins[i].Password, admins[i].FullName, admins[i].ShortName)
	}

	tableAdmin.AddFooters("", "", "", "", "Total Data:"+strconv.Itoa(len(admins)))
	tableAdmin.Render()
}
