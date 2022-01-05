package request

type (
	Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Register struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
	}
)
