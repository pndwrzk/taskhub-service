package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pndwrzk/taskhub-service/config"
	errConst "github.com/pndwrzk/taskhub-service/internal/constants/error"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID string) (string, time.Time, error) {
	exp := time.Now().Add(time.Duration(config.App.JWTAccessExpired) * time.Second)

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(config.App.JWTAccessSecret))
	return signed, exp, err
}

func GenerateRefreshToken(userID string) (string, time.Time, error) {
	exp := time.Now().Add(time.Duration(config.App.JWTRefreshExpired) * time.Second)

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(config.App.JWTRefreshSecret))
	return signed, exp, err
}

// Validate Refresh Token
func ValidateRefreshToken(tokenStr string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.App.JWTRefreshSecret), nil
	})
	if err != nil || !token.Valid {
		return "", errConst.ErrInvalidToken
	}
	return claims.UserID, nil
}
