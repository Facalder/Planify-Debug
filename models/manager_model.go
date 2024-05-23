package models

import (
	global "github.com/Facalder/Planify"
	"time"
)

type Manager_Model struct {
	Id         int
	FullName   string
	ShortName  string
	UserName   string
	Password   string
	Bio        string
	Department string
	CreatedAt  time.Time
	UpdateAt   time.Time
	DeletedAt  time.Time
}

type Tab_Manager [global.NMAX]Manager_Model

var (
	NManager int = global.NArray
	Managers Tab_Manager
)
