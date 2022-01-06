package repository

import "github.com/persiangophers/userland/entity"

type (
	Cash interface {
		// User
		SetUser(user entity.User) error
		GetUser(uuid string) (entity.User, error)
	}

	DB interface {
		// UUID
		CreateUUID(uuid entity.UUID) error
		GetUUIDByID(id uint) (entity.UUID, error)
		GetUUID(uuid string) (entity.UUID, error)
		DeleteUUID(id uint) error
		HardDeleteUUID(id uint) error

		// Username
		CreateUsername(username entity.Username) error
		GetUsernameByUUID(uuid string) (entity.Username, error)
		GetUsernameByID(id uint) (entity.Username, error)
		GetUsername(username string) (entity.Username, error)
		DeleteUsername(id uint) error
		HardDeleteUsername(id uint) error

		// Phone
		CreatePhone(phone entity.Phone) error
		GetPhoneByUUID(uuid string) (entity.Phone, error)
		GetPhoneByID(id uint) (entity.Phone, error)
		GetPhone(phoneNumber string) (entity.Phone, error)
		UpdatePhone(phone entity.Phone) error
		DeletePhone(id uint) error
		HardDeletePhone(id uint) error

		// Email
		CreateEmail(email entity.Email) error
		GetEmailByUUID(uuid string) (entity.Email, error)
		GetEmailByID(id uint) (entity.Email, error)
		GetEmail(email string) (entity.Email, error)
		UpdateEmail(email entity.Email) error
		DeleteEmail(id uint) error
		HardDeleteEmail(id uint) error

		// Profile
		CreateProfile(profile entity.Profile) error
		GetProfileByUUID(uuid string) (entity.Profile, error)
		GetProfileByID(id uint) (entity.Profile, error)
		UpdateProfile(profile entity.Profile) error
		DeleteProfile(id uint) error
		HardDeleteProfile(id uint) error

		// Password
		CreatePassword(password entity.Password) error
		GetPasswordByUUID(uuid string) (entity.Password, error)
		GetPassword(uuid uint) (entity.Password, error)
		DeletePassword(id uint) error
		HardDeletePassword(id uint) error
	}
)