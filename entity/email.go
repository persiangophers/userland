package entity

import "time"

type Email struct {
	ID         UUID
	Email      string
	VerifiedAt *time.Time
}
