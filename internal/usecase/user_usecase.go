// internal/usecase/user_usecase.go
package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/pndwrzk/taskhub-service/internal/common/utils"
	"github.com/pndwrzk/taskhub-service/internal/constants"
	errConst "github.com/pndwrzk/taskhub-service/internal/constants/error"
	"github.com/pndwrzk/taskhub-service/internal/dto"
	"github.com/pndwrzk/taskhub-service/internal/model"
	"github.com/pndwrzk/taskhub-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
	RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {

	existingUser, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, errConst.ErrUserNotFound) {
		return nil, err
	}

	if existingUser != nil {

		return nil, errConst.ErrEmailExist
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:       uuid.New(),
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (u *userUsecase) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errConst.ErrLogin
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errConst.ErrLogin
	}

	idStr := user.ID.String()

	accessToken, accessExp, err := utils.GenerateAccessToken(idStr)
	if err != nil {
		return nil, err
	}
	refreshToken, refreshExp, err := utils.GenerateRefreshToken(idStr)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessExp.Unix(),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshExp.Unix(),
		TokenType:             constants.TokenType,
		User: dto.UserInfo{
			ID:    user.ID,
			Email: user.Email,
		},
	}, nil
}

func (u *userUsecase) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {

	userID, err := utils.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	accessToken, accessExp, err := utils.GenerateAccessToken(userID)
	if err != nil {
		return nil, errConst.ErrInvalidToken
	}
	newRefreshToken, refreshExp, err := utils.GenerateRefreshToken(userID)
	if err != nil {
		return nil, errConst.ErrInvalidToken
	}

	return &dto.RefreshTokenResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessExp.Unix(),
		RefreshToken:          newRefreshToken,
		RefreshTokenExpiresAt: refreshExp.Unix(),
		TokenType:             constants.TokenType,
	}, nil
}
