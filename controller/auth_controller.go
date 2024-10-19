package controller

import (
	"context"
	"net/http"

	"github.com/elangreza14/moviefestival/dto"
	"github.com/gin-gonic/gin"
)

type (
	authService interface {
		RegisterUser(ctx context.Context, req dto.RegisterPayload) error
		LoginUser(ctx context.Context, req dto.LoginPayload) (string, error)
	}

	AuthController struct {
		authService authService
	}
)

func NewAuthController(authService authService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.RegisterPayload{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		err = ac.authService.RegisterUser(c, req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusCreated, dto.NewBaseResponse("created", nil))
	}
}

func (ac *AuthController) LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := dto.LoginPayload{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		token, err := ac.authService.LoginUser(c, req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse(token, nil))
	}
}
