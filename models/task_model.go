package models

import (
	global "github.com/Facalder/Planify"
	"time"
)

type Task_Model struct {
	Id          int
	Name        string
	Description string
	Priority    string  // low, med, high
	Status      string  // completed, not-completed
	Cost        float64 // task expected cost
	Notes       string
	Category    string // task category
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Tab_Task [global.NMAX]Task_Model

var (
	NTask int = global.NArray
	Tasks Tab_Task
)
