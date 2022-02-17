package protocol

import "time"

// Username is interface of Username entity
type (
	Username interface {
		UserID() [16]byte
		Username() string

		Status() UsernameStatus
	}

	UsernameStatus  uint8
	UsernameService interface {
		Create(username Username) error
		Find(username string) (Username, error)
		Last(userID [16]byte) (Username, error)
		Lasts(userID [16]byte, count int) ([]Username, error)
		Meantime(userID [16]byte, start, end time.Time) ([]Username, error)
		CountVersionUntil(userID [16]byte, until time.Time) (int, error)
	}
)

const (
	StatusUsernameUnSet UsernameStatus = 0b00000000
	StatusUsernameEmpty UsernameStatus = 0b00000001
)

type (
	InsertUsernameRequest interface {
		Username
	}

	GetUsernameRequest interface {
		UserID() [16]byte
	}

	GetUsernameResponse interface {
		Username
	}

	IsUsernameExistRequest interface {
		Username
	}

	IsUsernameExistResponse interface {
		Exist() bool
	}

	BookUsernameRequest interface {
		Username() string
	}

	GetUsernamesHistoryPerUserIDRequest interface {
		UserID() [16]byte
		Count() int
		StartTime() time.Time
		EndTime() time.Time
	}

	GetUsernamesHistoryPerUserIDResponse interface {
		Usernames() []Username
	}
)
