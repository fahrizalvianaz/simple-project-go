package users

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(ctx context.Context, user *User) (*User, error)
	FindUserByUsername(ctx context.Context, username string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Register(ctx context.Context, user *User) (*User, error) {
	result := r.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *userRepository) FindUserByUsername(ctx context.Context, username string) (*User, error) {
	var user *User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
