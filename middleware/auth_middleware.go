package middleware

import (
	"context"
	"errors"
	"fmt"
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

func (am *AuthMiddleware) MustHavePermissionMiddleware(permissions ...model.UserPermission) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRaw, ok := c.Get(UserMiddlewareKey)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.NewBaseResponse(nil, errors.New("cannot unauthorize this user")))
			return
		}

		user, ok := userRaw.(*model.User)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.NewBaseResponse(nil, errors.New("not valid user")))
			return
		}

		for _, permission := range permissions {
			ok := user.ValidPermission(permission.Val)
			if !ok {
				c.AbortWithStatusJSON(http.StatusForbidden, dto.NewBaseResponse(nil, fmt.Errorf("user doesn't have %s permission", permission.Name)))
				return
			}
		}

		c.Next()
	}
}
