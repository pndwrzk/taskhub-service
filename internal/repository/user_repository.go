// internal/repository/user_repository.go
package repository

import (
	"context"
	"errors"

	errConst "github.com/pndwrzk/taskhub-service/internal/constants/error"
	"github.com/pndwrzk/taskhub-service/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return errConst.ErrSqlError
	}

	return nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConst.ErrUserNotFound
		}
		return nil, errConst.ErrSqlError
	}

	return &user, nil
}
