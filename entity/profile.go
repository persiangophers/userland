package entity

import "time"

type Profile struct {
	ID        int
	UUID      string
	Firstname string
	Lastname  string
	BirthDate time.Time
	AboutMe   string
	Gender    string
	Links     []string
	Avatar    string
}
