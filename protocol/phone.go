package protocol

// Phone is interface of phone entity
type Phone interface {
	UserID() [16]byte
	Phone(userID [16]byte) string

	// list of status :
	// 0 = inactive
	// 1 = active
	Status() uint8
}

// PhoneService is interface of Phone storage layer
type PhoneService interface {
	Create(password Phone) error
	Get(id [16]byte) (Phone, error)
}

// Phone business layer
type (
	InsertPhoneRequest interface {
		Phone
	}

	GetPhoneRequest interface {
		UserID() UserID
	}

	GetPhoneResponse interface {
		Phone
	}
)
