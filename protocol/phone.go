package protocol

import "time"

// Phone is interface of phone entity
type (
	Phone interface {
		UserID() [16]byte
		Phone(userID [16]byte) string

		Status() PhoneStatus
	}

	PhoneStatus uint8

	PhoneService interface {
		Create(phone Phone) error
		Find(phone string) (Phone, error)
		Last(userID [16]byte) (Phone, error)
		Lasts(userID [16]byte, count int) ([]Phone, error)
		Meantime(userID [16]byte, start, end time.Time) ([]Phone, error)
		CountVersionUntil(userID [16]byte, until time.Time) (int, error)
	}
)

const (
	StatusPhoneUnSet    PhoneStatus = 0b00000000
	StatusPhoneEmpty    PhoneStatus = 0b00000001
	StatusPhoneInactive PhoneStatus = 0b00000010
	StatusPhoneActive   PhoneStatus = 0b00000100
)

// Phone business layer
type (
	InsertPhoneRequest interface {
		Phone() string
	}

	GetPhoneRequest interface {
		UserID() [16]byte
	}

	GetPhoneResponse interface {
		Phone
	}

	IsPhoneVerifiedRequest interface {
		Phone
	}

	IsPhoneVerifiedResponse interface {
		Verified() bool
	}

	SendPhoneVerificationRequest interface {
		Phone
	}

	IsPhoneExistRequest interface {
		Phone
	}

	IsPhoneExistResponse interface {
		Exist() bool
	}

	GetPhonesHistoryPerUserIDRequest interface {
		UserID() [16]byte
		Count() int
		StartTime() time.Time
		EndTime() time.Time
	}

	GetPhonesHistoryPerUserIDResponse interface {
		Phones() []Phone
	}
)
