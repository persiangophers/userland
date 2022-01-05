package entity

import "time"

type (
	User struct {
		ID            int        `json:"id"`
		UUID          string     `json:"uuid"`
		Firstname     string     `json:"firstname"`
		Lastname      string     `json:"lastname"`
		Picture       string     `json:"picture"`
		Username      string     `json:"username"`
		Email         string     `json:"email"`
		VerifiedEmail bool       `json:"verified_email"`
		Phonenumber   string     `json:"phonenumber"`
		VerifiedPhone bool       `json:"verified_phone"`
		Password      string     `json:"-"`
		UpdatedAt     time.Time  `json:"updated_at"`
		CreatedAt     time.Time  `json:"created_at"`
		DeletedAt     *time.Time `json:"deleted_at"`
	}
)
