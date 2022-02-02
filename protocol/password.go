package protocol

// Password is interface of password entity
type Password interface {
	UserID() [16]byte
	Password() string
}

// PasswordService is interface of Password storage layer
type PasswordService interface {
	GetPassword(userID [16]byte) (Password, error)
	CreatePassword(password Password) error
}

// Password business layer
type (
	InsertPasswordRequest interface {
		Password
	}

	InsertPasswordResponse interface {
		error
	}

	CheckPasswordRequest interface {
		Password
	}

	CheckPasswordResponse interface {
		Valid() bool
		error
	}
)
