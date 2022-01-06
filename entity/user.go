package entity

import "time"

type (
	User struct {
		UUID      string
		Username  string
		Email     string
		Phone     string
		Firstname string
		Lastname  string
		BirthDate time.Time
		AboutMe   string
		Gender    string
		Links     []string
		Avatar    string
	}
)
