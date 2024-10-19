package service

import (
	"context"

	"github.com/elangreza14/moviefestival/dto"
	"github.com/elangreza14/moviefestival/model"
)

type (
	movieRepo interface {
		GetAll(ctx context.Context) ([]model.Movie, error)
		Get(ctx context.Context, by string, val any, columns ...string) (*model.Movie, error)
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
			ID:           movie.ID,
			DeviceName:   movie.Name,
			Manufacturer: movie.Manufacturer,
			Price:        movie.Price,
			Image:        movie.Image,
			Stock:        movie.Stock,
		})
	}

	return res, nil
}
