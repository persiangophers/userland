package entity

import "time"

type Email struct {
	ID         uint
	UUID       string
	Value      string
	VerifiedAt *time.Time
	CreatedAt  time.Time
	DeletedAt  *time.Time
}
