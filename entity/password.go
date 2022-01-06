package entity

import "time"

type Password struct {
	ID        int
	UUID      string
	Value     string
	CreatedAt time.Time
	DeletedAt *time.Time
}
