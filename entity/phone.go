package entity

import "time"

type Phone struct {
	ID         uint
	UUID       string
	Number     string
	VerifiedAt *time.Time
	CreatedAt  time.Time
	DeletedAt  *time.Time
}
