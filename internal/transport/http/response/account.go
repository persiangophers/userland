package response

type (
	Login struct {
		Token string `json:"token"`
	}

	Error struct {
		Error string `json:"error"`
	}
)
