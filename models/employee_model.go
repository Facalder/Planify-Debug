package models

import (
	global "github.com/Facalder/Planify"
	"time"
)

type Employee_Model struct {
	Id         int
	FullName   string
	ShortName  string
	UserName   string
	Password   string
	Bio        string
	Department string
	Manager    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type Tab_Employee [global.NMAX]Employee_Model

var (
	NEmployee int = global.NArray
	Employees Tab_Employee
)
