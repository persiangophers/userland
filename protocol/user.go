package protocol

type UserID [16]byte

type User interface {
	ID() UserID
	Email() string
	Username() string
	Phone() string
	Firstname() string
	Lastname() string
	DisplayName() string
	Gender() string
	BirthDate() [8]uint
	Bio() string
	Avatar() string
}

type UserService interface {
	GetUser(userID [16]byte) (User, error)
	SetUser(user User) error
}