package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/elangreza14/moviefestival/dto"
	"github.com/gin-gonic/gin"
)

type (
	movieService interface {
		MovieList(ctx context.Context) (dto.MovieListResponse, error)
		CreateMovie(ctx context.Context, req dto.CreateMoviePayload) error
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

func (cc *MovieController) UploadMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		contentType, ok := file.Header["Content-Type"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("must be valid mp4")))
			return
		}

		if len(contentType) == 0 || contentType[0] != "video/mp4" {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, errors.New("filetype must be mp4")))
			return
		}

		dst := "public/" + file.Filename
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		// link := "/api/movies/" + dst
		c.JSON(http.StatusOK, dto.NewBaseResponse(dst, nil))
	}
}

func (cc *MovieController) CreateMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateMoviePayload
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		err = cc.movieService.CreateMovie(c, req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse("ok", nil))
	}
}
