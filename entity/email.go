package entity

import "time"

type Email struct {
	UUID       [16]byte
	Value      string
	VerifiedAt *time.Time
	CreatedAt  time.Time
	// TODO i think it helps for performance.
	//DeletedAt *time.Time
}
