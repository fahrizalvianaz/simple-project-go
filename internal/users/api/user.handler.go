package api

import (
	"bookstore-framework/internal/users"
	"bookstore-framework/internal/users/api/dto"
	"bookstore-framework/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService users.UserService
}

func NewUserHandler(userService users.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) RegisterHandler(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		pkg.BadRequestResponse(ctx, "Invalid Request format", err.Error())
	}

	response, err := h.userService.Register(ctx.Request.Context(), req)
	if err != nil {
		pkg.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	pkg.CreatedResponse(ctx, "User registered successfully", response)
}
