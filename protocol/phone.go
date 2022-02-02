package protocol

// Phone is interface of phone entity
type Phone interface {
	UserID() [16]byte
	Phone(userID [16]byte) string
	Verified() bool
}

// PhoneService is interface of Phone storage layer
type PhoneService interface {
	Create(password Phone) error
	Get(id [16]byte) (Phone, error)
	GetVerified(id [16]byte) (bool, error)
}

// Phone business layer
type (
	InsertPhoneRequest interface {
		Phone
	}

	InsertPhoneResponse interface {
		error
	}

	GetPhoneRequest interface {
		UserID() UserID
	}

	GetPhoneResponse interface {
		Phone
		error
	}
)
