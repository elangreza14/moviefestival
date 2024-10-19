package controller

import (
	"context"
	"net/http"

	"github.com/elangreza14/moviefestival/dto"
	"github.com/gin-gonic/gin"
)

type (
	movieService interface {
		MovieList(ctx context.Context) (dto.MovieListResponse, error)
	}

	MovieController struct {
		movieService
	}
)

func NewMovieController(movieService movieService) *MovieController {
	return &MovieController{movieService}
}

func (cc *MovieController) MovieList() gin.HandlerFunc {
	return func(c *gin.Context) {
		currencies, err := cc.movieService.MovieList(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse(currencies, nil))
	}
}
