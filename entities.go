package main

import (
	"time"
)

type Alert struct {
	id           float64
	creationDate time.Time
	message      string
}

type Elder struct {
	id        string
	name      string
	lastName  string
	cellphone int64
	email     string
	relatives []string
}

type Relative struct {
	id                    string
	relationshipWithElder string // father, mother, bro, sis, friend, etc.
	name                  string
	lastName              string
	cellphone             int64
	email                 string
}
