package entity

import "time"

type User struct {
	ID          UUID
	Username    string
	Email       string
	Phone       string
	Firstname   string
	Lastname    string
	DisplayName string
	BirthDate   time.Time
	AboutMe     string
	Gender      string
	Links       []string
	Avatar      string
}
