package entity

import "time"

type Password struct {
	UUID      [16]byte
	Value     string
	CreatedAt time.Time
}
