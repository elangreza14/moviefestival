package service

//go:generate mockgen -source $GOFILE -destination ../../mock/service/mock_$GOFILE -package mock$GOPACKAGE

import (
	"context"
	"errors"
	"time"

	"github.com/elangreza14/moviefestival/internal/domain"
	"github.com/elangreza14/moviefestival/internal/dto"
	"github.com/jackc/pgx/v5"
)

type (
	movieRepo interface {
		GetAll(ctx context.Context) ([]domain.Movie, error)
		Get(ctx context.Context, by string, val any, columns ...string) (*domain.Movie, error)
		CreateMovieTX(ctx context.Context, movie domain.Movie, artists []string, genres []string) error
		UpdateMovieTX(ctx context.Context, movie domain.Movie, artists []string, genres []string) error
		GetMovieDetail(ctx context.Context, movieID int) (*domain.Movie, error)
		GetMoviesWithPaginationAndSearch(ctx context.Context, search, searchBy, orderBy, orderDirection string, page, pageSize int) ([]domain.Movie, error)
	}

	movieViewRepo interface {
		AddMovieViewTX(ctx context.Context, movieID int) error
	}

	movieService struct {
		movieRepo     movieRepo
		movieViewRepo movieViewRepo
		config        ServiceConfig
	}
)

func NewMovieService(movieRepo movieRepo, movieViewRepo movieViewRepo, config ServiceConfig) *movieService {
	return &movieService{
		movieRepo:     movieRepo,
		movieViewRepo: movieViewRepo,
		config:        config,
	}
}

func (cs *movieService) MovieList(ctx context.Context, req dto.MovieListParams) (dto.MovieListResponse, error) {
	movie, err := cs.movieRepo.GetMoviesWithPaginationAndSearch(ctx, req.Search, req.SearchBy, req.OrderBy, req.OrderDirection, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	res := make([]dto.MovieListResponseElement, 0)
	for _, movie := range movie {
		res = append(res, dto.MovieListResponseElement{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			Duration:    movie.Duration.String(),
			Views:       movie.Views,
			Artists:     movie.Artist,
			Genres:      movie.Genres,
			WatchUrl:    cs.config.BaseURL + "/api/movies/public/" + movie.WatchUrl,
		})
	}

	return res, nil
}

func (cs *movieService) CreateMovie(ctx context.Context, req dto.CreateMoviePayload) error {
	duration, err := time.ParseDuration(req.Duration)
	if err != nil {
		return errors.New("cannot parse duration")
	}

	movie := domain.NewMovie(req.Title, req.Description, req.WatchUrl, duration)
	return cs.movieRepo.CreateMovieTX(ctx, *movie, req.Artists, req.Genres)
}

func (cs *movieService) UpdateMovie(ctx context.Context, req dto.CreateMoviePayload, movieID int) error {
	movie, err := cs.movieRepo.Get(ctx, "id", movieID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.ErrorNotFound{Entity: "movie"}
		}
		return err
	}

	duration, err := time.ParseDuration(req.Duration)
	if err != nil {
		return errors.New("cannot parse duration")
	}
	movie.Duration = duration
	movie.Title = req.Title
	movie.Description = req.Description
	movie.WatchUrl = req.WatchUrl

	return cs.movieRepo.UpdateMovieTX(ctx, *movie, req.Artists, req.Genres)
}

func (cs *movieService) GetMovieDetail(ctx context.Context, movieID int) (*dto.MovieListResponseElement, error) {

	movie, err := cs.movieRepo.GetMovieDetail(ctx, movieID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, dto.ErrorNotFound{Entity: "movie"}
		}
		return nil, err
	}

	if err = cs.movieViewRepo.AddMovieViewTX(ctx, movieID); err != nil {
		return nil, err
	}

	return &dto.MovieListResponseElement{
		ID:          movieID,
		Title:       movie.Title,
		Description: movie.Description,
		Duration:    movie.Duration.String(),
		WatchUrl:    cs.config.BaseURL + "/api/movies/public/" + movie.WatchUrl,
		Views:       movie.Views,
		Artists:     movie.Artist,
		Genres:      movie.Genres,
	}, nil
}

func (cs *movieService) MostViewedMovieList(ctx context.Context) (dto.MovieListResponse, error) {
	movie, err := cs.movieRepo.GetMoviesWithPaginationAndSearch(ctx, "", "", "views", "desc", 1, 10)
	if err != nil {
		return nil, err
	}

	res := make([]dto.MovieListResponseElement, 0)
	for _, movie := range movie {
		res = append(res, dto.MovieListResponseElement{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			Duration:    movie.Duration.String(),
			Views:       movie.Views,
			Artists:     movie.Artist,
			Genres:      movie.Genres,
			WatchUrl:    movie.WatchUrl,
		})
	}

	return res, nil
}
