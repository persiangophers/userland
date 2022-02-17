package protocol

import "time"

type (
	// Email is interface of email entity
	Email interface {
		UserID() [16]byte
		Email() string

		Status() EmailStatus
	}

	EmailStatus uint8

	// EmailService is interface of Email storage layer
	EmailService interface {
		Create(email Email) error
		Find(email string) (Email, error)
		Last(userID [16]byte) (Email, error)
		Lasts(userID [16]byte, count int) ([]Email, error)
		Meantime(userID [16]byte, start, end time.Time) ([]Email, error)
		CountVersionUntil(userID [16]byte, until time.Time) (int, error)
	}
)

const (
	StatusEmailUnSet    EmailStatus = 0b00000000
	StatusEmailEmpty    EmailStatus = 0b00000001
	StatusEmailInactive EmailStatus = 0b00000010
	StatusEmailActive   EmailStatus = 0b00000100
)

// Email business layer
type (
	InsertEmailRequest interface {
		Email
	}

	GetEmailRequest interface {
		UserID() [16]byte
	}

	GetEmailResponse interface {
		Email
	}

	IsEmailVerifiedRequest interface {
		Email
	}

	IsEmailVerifiedResponse interface {
		Verified() bool
	}

	SendEmailVerificationRequest interface {
		Email
	}

	IsEmailExistRequest interface {
		Email
	}

	IsEmailExistResponse interface {
		Exist() bool
	}

	GetEmailsHistoryPerUserIDRequest interface {
		UserID() [16]byte
		Count() int
		StartTime() time.Time
		EndTime() time.Time
	}

	GetEmailsHistoryPerUserIDResponse interface {
		Emails() []Email
	}
)
