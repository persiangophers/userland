package protocol

// Username is interface of Username entity
type Username interface {
	UserID() [16]byte
	Username() string
}

// UsernameService is interface of Username storage layer
type UsernameService interface {
	CreateUsername(username Username) error
	GetUsername(userID [16]byte) (Username, error)
}

// Username business layer
type (
	InsertUsernameRequest interface {
		Username
	}

	InsertUsernameResponse interface {
		error
	}

	GetUsernameRequest interface {
		UserID() UserID
	}

	GetUsernameResponse interface {
		Username
		error
	}
)
