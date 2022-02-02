package protocol

type (
	// Email is interface of email entity
	Email interface {
		UserID() [16]byte
		Email() string
		Verified() bool
	}

	// EmailService is interface of Email storage layer
	EmailService interface {
		Create(email Email) error
		Get(id [16]byte) (Email, error)
		GetVerified(id [16]byte) (bool, error)
	}
)

const GetEmailURIPath string = "api/v1/get/email"

// Email business layer
type (
	InsertEmailRequest interface {
		Email
	}

	InsertEmailResponse interface {
		error
	}

	GetEmailRequest interface {
		UserID() UserID
	}

	GetEmailResponse interface {
		Email
		error
	}
)
