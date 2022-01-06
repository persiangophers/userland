package entity

import "time"

type Phone struct {
	ID         int
	UUID       string
	Number     string
	VerifiedAt *time.Time
	CreatedAt  time.Time
	DeletedAt  *time.Time
}
