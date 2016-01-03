package panicdevice

import (
	"time"
)

type PanicDevice struct {
	serial    string
	birthDate time.Time
	beathDate time.Time
}
