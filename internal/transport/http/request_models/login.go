package request_models

type Login struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
