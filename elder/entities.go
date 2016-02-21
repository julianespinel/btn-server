package elder

import (
	"time"
)

type Elder struct {
	Id               string `json:"id" binding:"required"`
	Name             string `json:"name" binding:"required"`
	LastName         string `json:"lastName" binding:"required"`
	Cellphone        string
	RegistrationDate time.Time
}

type Relative struct {
	Id           string  `json:"id" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	LastName     string  `json:"lastName" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	Cellphone    string `json:"cellphone" binding:"required"`
	Relationship string  `json:"relationship" binding:"required"`
}
