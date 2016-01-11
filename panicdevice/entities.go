package panicdevice

import (
	"time"
)

type PanicDevice struct {
	Serial    string    `json:"serial" binding:"required"`
	BirthDate time.Time `json:"birthDate" binding:"required"`
	DeathDate time.Time
}
