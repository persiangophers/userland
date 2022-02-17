package protocol

import "time"

// Password is interface of password entity
type Password interface {
	UserID() [16]byte
	Password() string
}

// PasswordService is interface of Password storage layer
type PasswordService interface {
	Create(password Password) error
	Last(userID [16]byte) (Password, error)
	Lasts(userID [16]byte, count int) ([]Password, error)
	Meantime(userID [16]byte, start, end time.Time) ([]Password, error)
}

// Password business layer
type (
	InsertPasswordRequest interface {
		Password
	}

	CheckPasswordRequest interface {
		Password
	}

	CheckPasswordResponse interface {
		Valid() bool
	}

	GetPasswordsHistoryPerUserIDRequest interface {
		UserID() [16]byte
		Count() int
		StartTime() time.Time
		EndTime() time.Time
	}

	GetPasswordsHistoryPerUserIDResponse interface {
		Passwords() []Password
	}
)
