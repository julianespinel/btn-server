package elder

import (
	"time"
)

type Elder struct {
	Id               string `json:"id" binding:"required"`
	Name             string `json:"name" binding:"required"`
	LastName         string `json:"lastName" binding:"required"`
	Cellphone        float64
	RegistrationDate time.Time
}

type Relative struct {
	Id           string  `json:"id" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	LastName     string  `json:"lastName" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	Cellphone    float64 `json:"cellphone" binding:"required"`
	Relationship string  `json:"relationship" binding:"required"`
}
