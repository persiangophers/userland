package contract

import (
	"github.com/kianooshaz/userland/internal/dto"
	"github.com/kianooshaz/userland/internal/entity"
)

type (
	Ident interface {
		FindUser(identity string) (*entity.User, error)

		GetUserWithID(id int) (*entity.User, error)
		GetUserWithUUID(uuid string) (*entity.User, error)
		GetUserWithUsername(username string) (*entity.User, error)
		GetUserWithPhonenumber(phonenumber string) (*entity.User, error)
		GetUserWithEmail(email string) (*entity.User, error)
	}

	Auth interface {
		AddUser(req *dto.CreateUserRequest) error

		SendPhonenumberVerification(phonenumber string) error
		VerifyPhonenumber(phonenumber, code string) error

		SendEmailVerification(phonenumber string) error
		VerifyEmail(phonenumber, code string) error

		UpdateUsername(id int, username string) error
		UpdateEmail(id int, email string) error
		UpdatePhonenumber(id int, phonenumber string) error
		UpdateProfile(req *dto.UpdateProfileRequest) error

		Verify(identity, password string) error

		ForgetPassword(identity string) error
		ChangePassword(identity, code string) error
	}
)
