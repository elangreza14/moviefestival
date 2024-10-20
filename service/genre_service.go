package service

import (
	"context"

	"github.com/elangreza14/moviefestival/dto"
	"github.com/elangreza14/moviefestival/model"
)

type (
	genreRepo interface {
		GetMostViewedGenres(ctx context.Context) ([]model.ViewedGenre, error)
	}

	genreService struct {
		genreRepo genreRepo
	}
)

func NewGenreService(genreRepo genreRepo) *genreService {
	return &genreService{
		genreRepo: genreRepo,
	}
}

func (cs *genreService) MostViewedGenreList(ctx context.Context) (dto.GenreListResponse, error) {
	genres, err := cs.genreRepo.GetMostViewedGenres(ctx)
	if err != nil {
		return nil, err
	}

	res := dto.GenreListResponse{}
	for _, genre := range genres {
		res = append(res, dto.MostViewedGenreListResponseElement{
			Name:  genre.Name,
			Views: genre.Views,
		})
	}

	return res, nil
}
