package myjwt

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/golang-jwt/jwt/v4"
)

type myJWT struct {
	TokenClaims []string
	Secret      []byte
}

var instance *myJWT
var lock = &sync.Mutex{}

func GetInstance() Myjwt {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &myJWT{}
		}
	}
	return instance
}

func (m *myJWT) SetSecret(secret []byte) {
	m.Secret = make([]byte, len(secret))
	copy(m.Secret, secret)
}

func (m *myJWT) SetClaims(claims ...string) {
	m.TokenClaims = make([]string, len(claims))
	copy(m.TokenClaims, claims)
}

func (m *myJWT) Claims() []string {
	return m.TokenClaims
}

func (m *myJWT) NewToken(claims ...interface{}) (string, error) {

	var tokenClaims = jwt.MapClaims{}

	if len(claims) != len(m.TokenClaims) {
		return "", errors.New("claim's size is not right")
	}

	for i, claim := range claims {
		tokenClaims[m.TokenClaims[i]] = claim
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	if tokenString, err := token.SignedString(m.Secret); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func (m *myJWT) IsValid(tokenStr string) (bool, []interface{}) {
	var claimsValues = make([]interface{}, len(m.TokenClaims))
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return m.Secret, nil
	})

	if err != nil {
		return false, nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for key, value := range claims {
			for i, name := range m.TokenClaims {
				if name == key {
					claimsValues[i] = value
				}
			}
		}
	}
	return true, claimsValues
}

func FetchToken(bearer string) string {
	parts := strings.Split(bearer, " ")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}
