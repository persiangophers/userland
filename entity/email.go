package entity

import "time"

type Email struct {
	UUID       [16]byte
	Value      string
	VerifiedAt *time.Time
}
