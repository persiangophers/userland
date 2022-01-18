package entity

import "time"

type UUID struct {
	Value     [16]byte
	CreatedAt time.Time
}
