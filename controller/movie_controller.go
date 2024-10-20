package controller

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/elangreza14/moviefestival/dto"
	"github.com/gin-gonic/gin"
)

type (
	movieService interface {
		MovieList(ctx context.Context) (dto.MovieListResponse, error)
		CreateMovie(ctx context.Context, req dto.CreateMoviePayload) error
		UpdateMovie(ctx context.Context, req dto.CreateMoviePayload, movieID int) error
		GetMovieDetail(ctx context.Context, movieID int) (*dto.MovieListResponseElement, error)
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

func (cc *MovieController) UpdateMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateMoviePayload
		err := c.ShouldBindBodyWithJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		movieIDRaw := c.Param("movieID")
		movieID, err := strconv.Atoi(movieIDRaw)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		err = cc.movieService.UpdateMovie(c, req, movieID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse("ok", nil))
	}
}

func (cc *MovieController) GetMovieDetail() gin.HandlerFunc {
	return func(c *gin.Context) {

		movieIDRaw := c.Param("movieID")
		movieID, err := strconv.Atoi(movieIDRaw)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		movie, err := cc.movieService.GetMovieDetail(c, movieID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusOK, dto.NewBaseResponse(movie, nil))
	}
}
