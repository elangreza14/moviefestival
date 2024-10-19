package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/elangreza14/moviefestival/dto"
	"github.com/elangreza14/moviefestival/model"
	"github.com/gin-gonic/gin"
)

type (
	authService interface {
		ProcessToken(ctx context.Context, reqToken string) (*model.User, error)
	}

	AuthMiddleware struct {
		authService
	}
)

func NewAuthMiddleware(AuthService authService) *AuthMiddleware {
	return &AuthMiddleware{AuthService}
}

const UserMiddlewareKey = "UserMiddlewareKey"

func (am *AuthMiddleware) MustAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawAuthorization := c.Request.Header["Authorization"]
		if len(rawAuthorization) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("token not valid")))
			return
		}

		authorization := c.Request.Header["Authorization"][0]

		rawToken := strings.Split(authorization, " ")
		if len(rawToken) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("token not valid")))
			return
		}

		token := rawToken[1]

		user, err := am.authService.ProcessToken(c, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.NewBaseResponse(nil, errors.New("cannot unauthorize this user")))
			return
		}

		c.Set(UserMiddlewareKey, user)

		c.Next()
	}
}
