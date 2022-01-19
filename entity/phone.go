package entity

import "time"

type Phone struct {
	ID         UUID
	Code       int
	Number     int
	VerifiedAt *time.Time
}
