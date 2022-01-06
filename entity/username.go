package entity

import "time"

type Username struct {
	ID        int
	UUID      string
	Value     string
	CreatedAt time.Time
	DeletedAt *time.Time
}
