package service_test

import (
	"bookstore-framework/internal/users"
	"bookstore-framework/internal/users/api/dto"
	mocks "bookstore-framework/test/mock"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Success(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	service := users.NewUserService(mockRepo)

	ctx := context.Background()
	req := dto.RegisterRequest{
		Username: "test",
		Name:     "testuser",
		Email:    "test@gmail.com",
		Password: "password123",
	}

	expectedUser := &users.User{
		ID:       1,
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
	}

	mockRepo.EXPECT().Register(gomock.Any(), gomock.Any()).Return(expectedUser, nil)

	result, err := service.Register(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, result.ID)
	assert.Equal(t, req.Username, result.Username)
}
