package handler_test

import (
	"bookstore-framework/internal/users/api"
	"bookstore-framework/internal/users/api/dto"
	mocks "bookstore-framework/test/mock"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockUserService(ctrl)
	handler := api.NewUserHandler(mockService)

	req := dto.RegisterRequest{
		Username: "test",
		Name:     "testuser",
		Email:    "test@gmail.com",
		Password: "password123",
	}

	data := dto.RegisterResponse{
		ID:        1,
		Username:  req.Username,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	mockService.EXPECT().Register(gomock.Any(), gomock.Eq(req)).
		Return(&data, nil)

	body, err := json.Marshal(req)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.RegisterHandler(c)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response dto.RegisterResponse
	err = json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, data.ID, data.ID)
	assert.Equal(t, data.Username, response.Username)

}
