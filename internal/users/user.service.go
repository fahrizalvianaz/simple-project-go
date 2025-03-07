package users

import (
	"bookstore-framework/internal/users/api/dto"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(ctx context.Context, req dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := User{
		Username: req.Username,
		Name:     req.Name,
		Password: string(hashedPassword),
		Email:    req.Email,
	}

	registerUser, err := s.userRepo.Register(ctx, &user)
	if err != nil {
		return nil, err
	}

	respone := &dto.RegisterResponse{
		ID:        registerUser.ID,
		Username:  registerUser.Username,
		Email:     registerUser.Email,
		CreatedAt: registerUser.CreatedAt,
	}

	return respone, nil

}

func (s *userService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid password or username")
	}

	respose := &dto.LoginResponse{
		TokenAccess: "ini acess token",
	}

	return respose, nil

}
