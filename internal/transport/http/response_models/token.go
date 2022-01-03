package response_models

type Token struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
