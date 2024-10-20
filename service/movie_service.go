package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/elangreza14/moviefestival/dto"
	"github.com/elangreza14/moviefestival/model"
	"github.com/jackc/pgx/v5"
)

type (
	movieRepo interface {
		GetAll(ctx context.Context) ([]model.Movie, error)
		Get(ctx context.Context, by string, val any, columns ...string) (*model.Movie, error)
		CreateMovieTX(ctx context.Context, movie model.Movie, artists []string, genres []string) error
		UpdateMovieTX(ctx context.Context, movie model.Movie, artists []string, genres []string) error
		GetMovieDetail(ctx context.Context, movieID int) (*model.Movie, error)
		GetMoviesWithPaginationAndSearch(ctx context.Context, search, searchBy, orderBy, orderDirection string, page, pageSize int) ([]model.Movie, error)
	}

	movieViewRepo interface {
		AddMovieViewTX(ctx context.Context, movieID int) error
	}

	movieService struct {
		movieRepo     movieRepo
		movieViewRepo movieViewRepo
	}
)

func NewMovieService(movieRepo movieRepo, movieViewRepo movieViewRepo) *movieService {
	return &movieService{
		movieRepo:     movieRepo,
		movieViewRepo: movieViewRepo,
	}
}

func (cs *movieService) MovieList(ctx context.Context, req dto.MovieListParams) (dto.MovieListResponse, error) {
	movie, err := cs.movieRepo.GetMoviesWithPaginationAndSearch(ctx, req.Search, req.SearchBy, req.OrderBy, req.OrderDirection, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	baseUrl := fmt.Sprintf("http://localhost%s", os.Getenv("HTTP_PORT"))

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
			WatchUrl:    baseUrl + "/api/movies/public/" + movie.WatchUrl,
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

func (cs *movieService) UpdateMovie(ctx context.Context, req dto.CreateMoviePayload, movieID int) error {
	movie, err := cs.movieRepo.Get(ctx, "id", movieID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("movie not found")
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
	err := cs.movieViewRepo.AddMovieViewTX(ctx, movieID)
	if err != nil {
		return nil, err
	}

	movie, err := cs.movieRepo.GetMovieDetail(ctx, movieID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("movie not found")
		}
		return nil, err
	}

	baseUrl := fmt.Sprintf("http://localhost%s", os.Getenv("HTTP_PORT"))

	return &dto.MovieListResponseElement{
		ID:          movieID,
		Title:       movie.Title,
		Description: movie.Description,
		Duration:    movie.Duration.String(),
		WatchUrl:    baseUrl + "/api/movies/public/" + movie.WatchUrl,
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
