package protocol

type (
	// Email is interface of email entity
	Email interface {
		UserID() [16]byte
		Email() string

		// list of status :
		// 0 = inactive
		// 1 = active
		Status() uint8
	}

	// EmailService is interface of Email storage layer
	EmailService interface {
		Create(email Email) error
		Get(id [16]byte) (Email, error)
	}
)

const GetEmailURIPath string = "api/v1/get/email"

// Email business layer
type (
	InsertEmailRequest interface {
		Email
	}

	GetEmailRequest interface {
		UserID() UserID
	}

	GetEmailResponse interface {
		Email
	}
)
