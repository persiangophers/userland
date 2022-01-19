package entity

import "time"

type Profile struct {
	ID          UUID
	Firstname   string
	Lastname    string
	DisplayName string
	BirthDate   time.Time
	AboutMe     string
	Gender      string
	Links       []string
	Avatar      string
}
