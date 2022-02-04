package protocol

import "time"

// Username is interface of Username entity
type Username interface {
	UserID() [16]byte
	Username() string
}

// UsernameService is interface of Username storage layer
type UsernameService interface {
	// CreateUsername is used when want assign new username to userID
	CreateUsername(username Username) error
	// GetUsername is used when want get Username by username
	GetUsername(username string) (Username, error)
	// GetLastUsername is used when want get last Username of userID
	GetLastUsername(userID [16]byte) (Username, error)
	// GetLastUsernames is used when want the list of userID's last Usernames.
	// if count equal 0 it means there is no limit
	GetLastUsernames(userID [16]byte, count int) ([]Username, error)
	// GetLastUsernames is used when want the list of userID's Usernames between time intervals.
	GetUsernamesBetweenTime(start, end time.Time) ([]Username, error)
}

// Username business layer
type (
	InsertUsernameRequest interface {
		Username
	}

	GetUsernameRequest interface {
		UserID() UserID
	}

	GetUsernameResponse interface {
		Username
	}

	IsUsernameUniqueRequest interface {
		Username
	}
)
