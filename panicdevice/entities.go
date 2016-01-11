package panicdevice

import (
	"time"
)

type PanicDevice struct {
	Serial    string    `binding:"required"`
	BirthDate time.Time `binding:"required"`
	DeathDate time.Time
}
