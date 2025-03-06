package users

import (
	"bookstore-framework/internal/users/api/dto"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req dto.RegisterRequest) (*dto.RegisterResponse, error)
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error) {
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
	registerUser, err := s.userRepo.Register(&user)
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
