package entity

import "time"

type Username struct {
	UUID      [16]byte
	Value     string
	CreatedAt time.Time
}
