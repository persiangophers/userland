package entity

import "time"

type Email struct {
	ID         UUID
	Value      string
	VerifiedAt *time.Time
}
