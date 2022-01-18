package entity

import "time"

type Phone struct {
	UUID       [16]byte
	Number     string
	VerifiedAt *time.Time
	CreatedAt  time.Time
}
