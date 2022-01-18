package entity

import "time"

type Profile struct {
	UUID        [16]byte
	Firstname   string
	Lastname    string
	DisplayName string
	BirthDate   time.Time
	AboutMe     string
	Gender      string
	Links       []string
	Avatar      string
}
