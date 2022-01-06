package entity

import "time"

type Username struct {
	ID        uint
	UUID      string
	Value     string
	CreatedAt time.Time
	DeletedAt *time.Time
}
