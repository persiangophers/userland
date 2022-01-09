package entity

import "time"

type Password struct {
	ID        uint
	UUID      string
	Value     string
	CreatedAt time.Time
	DeletedAt *time.Time
}