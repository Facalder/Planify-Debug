package models

import (
	global "github.com/Facalder/Planify"
	"time"
)

type Admin_Model struct {
	Id        int
	FullName  string
	ShortName string
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}

type Tab_Admin [global.NMAX]Admin_Model

var (
	NAdmin int = global.NArray
	Admins Tab_Admin
)
