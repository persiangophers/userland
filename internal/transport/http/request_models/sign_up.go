package request_models

type SignUp struct {
	FirstName  string   `validate:"required,min=3,max=32" json:"first_name" form:"first_name"`
	LastName   string   `validate:"required,min=3,max=32" json:"last_name" form:"last_name"`
	Privileges []string `validate:"required,gt=0,dive,required"`
	Avatar     string
	Email      string `validate:"required,email,min=6,max=32"`
	Password   string `validate:"required,min=6,max=32"`
	BirthDate  string `validate:"required" json:"birth_date" form:"birth_date"`
}
