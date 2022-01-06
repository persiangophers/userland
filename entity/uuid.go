package entity

import "time"

type UUID struct {
	ID        uint
	Value     string
	CreatedAt time.Time
	DeletedAt *time.Time
}
