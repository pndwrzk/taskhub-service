package dto

import (
	"time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	AccessToken           string   `json:"access_token"`
	AccessTokenExpiresAt  int64    `json:"access_token_expires_at"`
	RefreshToken          string   `json:"refresh_token"`
	RefreshTokenExpiresAt int64    `json:"refresh_token_expires_at"`
	TokenType             string   `json:"token_type"`
	User                  UserInfo `json:"user"`
}

type UserInfo struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenResponse struct {
	AccessToken           string `json:"access_token"`
	AccessTokenExpiresAt  int64  `json:"access_token_expires_at"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
	TokenType             string `json:"token_type"`
}
