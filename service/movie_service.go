package service

import (
	"context"
	"errors"
	"time"

	"github.com/elangreza14/moviefestival/dto"
	"github.com/elangreza14/moviefestival/model"
)

type (
	movieRepo interface {
		GetAll(ctx context.Context) ([]model.Movie, error)
		Get(ctx context.Context, by string, val any, columns ...string) (*model.Movie, error)
		CreateMovieTX(ctx context.Context, movie model.Movie, artists []string, genres []string) error
	}

	movieService struct {
		movieRepo movieRepo
	}
)

func NewMovieService(movieRepo movieRepo) *movieService {
	return &movieService{
		movieRepo: movieRepo,
	}
}

func (cs *movieService) MovieList(ctx context.Context) (dto.MovieListResponse, error) {
	movie, err := cs.movieRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]dto.MovieListResponseElement, 0)
	for _, movie := range movie {
		res = append(res, dto.MovieListResponseElement{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			Duration:    movie.Duration,
			Artists:     []string{},
			Genres:      []string{},
			WatchUrl:    movie.WatchUrl,
		})
	}

	return res, nil
}

func (cs *movieService) CreateMovie(ctx context.Context, req dto.CreateMoviePayload) error {
	duration, err := time.ParseDuration(req.Duration)
	if err != nil {
		return errors.New("cannot parse duration")
	}

	movie := model.NewMovie(req.Title, req.Description, req.WatchUrl, duration)
	return cs.movieRepo.CreateMovieTX(ctx, *movie, req.Artists, req.Genres)
}
