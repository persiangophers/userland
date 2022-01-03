package jwt_generator

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/persiangophers/userland/internal/config"
	"github.com/persiangophers/userland/internal/transport/http/response_models"

	"time"
)

func GenerateToken(userId string, cfg *config.Config) (response_models.Token, error) {
	accessClaims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	access, err := accessToken.SignedString([]byte(cfg.Jwt.Secret))

	if err != nil {
		return response_models.Token{}, err
	}

	refreshClaims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	refresh, err := refreshToken.SignedString([]byte(cfg.Jwt.Secret))

	if err != nil {
		return response_models.Token{}, err
	}

	return response_models.Token{
		Access:  access,
		Refresh: refresh,
	}, nil
}
