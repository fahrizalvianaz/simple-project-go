package users

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(ctx context.Context, user *User) (*User, error)
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
