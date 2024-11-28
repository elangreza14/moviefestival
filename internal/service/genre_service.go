package service

//go:generate mockgen -source $GOFILE -destination ../../mock/service/mock_$GOFILE -package mock$GOPACKAGE

import (
	"context"

	"github.com/elangreza14/moviefestival/internal/domain"
	"github.com/elangreza14/moviefestival/internal/dto"
)

type (
	genreRepo interface {
		GetMostViewedGenres(ctx context.Context) ([]domain.ViewedGenre, error)
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
