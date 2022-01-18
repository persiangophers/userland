package entity

import "time"

type Phone struct {
	UUID       [16]byte
	Code       int
	Number     int
	VerifiedAt *time.Time
	CreatedAt  time.Time
}
