package entity

import "time"

type Email struct {
	ID         int
	UUID       string
	Value      string
	VerifiedAt *time.Time
	CreatedAt  time.Time
	DeletedAt  *time.Time
}
