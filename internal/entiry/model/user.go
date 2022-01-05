package model

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"-"`
	Email     string `json:"email"`
}
