package myjwt

type Myjwt interface {
	SetSecret([]byte)
	SetClaims(claims ...string)
	Claims() []string
	NewToken(claimsValue ...interface{}) (string, error)
	IsValid(tokenStr string) (bool, []interface{})
}
