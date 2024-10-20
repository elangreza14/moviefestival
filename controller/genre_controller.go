package controller

import (
	"context"
	"net/http"

	"github.com/elangreza14/moviefestival/dto"
	"github.com/gin-gonic/gin"
)

type (
	genreService interface {
		MostViewedGenreList(ctx context.Context) (dto.GenreListResponse, error)
	}

	GenreController struct {
		genreService
	}
)

func NewGenreController(genreService genreService) *GenreController {
	return &GenreController{genreService}
}

func (cc *GenreController) MostViewedGenreList() gin.HandlerFunc {
	return func(c *gin.Context) {

		res, err := cc.genreService.MostViewedGenreList(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse(res, nil))
	}
}
