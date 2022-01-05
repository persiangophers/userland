package dto

type (
	CreateUserRequest struct {
		Firstname   string `json:"firstname"`
		Lastname    string `json:"lastname"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		Phonenumber string `json:"phonenumber"`
		Password    string `json:"password"`
	}

	UpdateProfileRequest struct {
		ID        int    `json:"id"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Picture   string `json:"picture"`
	}
)
