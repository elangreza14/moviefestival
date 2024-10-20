package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/elangreza14/moviefestival/model"
)

type movieRepository struct {
	db QueryPgx
	*PostgresRepo[model.Movie]
	txRepo *PostgresTransactionRepo
}

func NewMovieRepository(
	dbPool QueryPgx,
	tx PgxTXer,
) *movieRepository {
	return &movieRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Movie](dbPool),
		txRepo:       NewPostgresTransactionRepo(tx),
	}
}

func (ur *movieRepository) CreateMovieTX(ctx context.Context, movie model.Movie, artists []string, genres []string) error {
	return ur.txRepo.WithTX(ctx, func(tx QueryPgx) error {
		movieRepo := NewMovieRepository(tx, nil)
		err := movieRepo.CreateAndScanId(ctx, &movie)
		if err != nil {
			return err
		}

		artistRepo := NewArtistRepository(tx)
		for _, artist := range artists {
			err = artistRepo.UpsertArtist(ctx, artist)
			if err != nil {
				return err
			}

			err = artistRepo.InsertArtistWithMovieMapping(ctx, artist, movie.ID)
			if err != nil {
				return err
			}
		}

		genreRepo := NewGenreRepository(tx)
		for _, genre := range genres {
			err = genreRepo.UpsertGenre(ctx, genre)
			if err != nil {
				return err
			}

			err = genreRepo.InsertGenreWithMovieMapping(ctx, genre, movie.ID)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (ur *movieRepository) UpdateMovieTX(ctx context.Context, movie model.Movie, artists []string, genres []string) error {
	return ur.txRepo.WithTX(ctx, func(tx QueryPgx) error {
		movieRepo := NewMovieRepository(tx, nil)
		whereValues := map[string]any{}
		whereValues["id"] = movie.ID
		err := movieRepo.Edit(ctx, movie, whereValues)
		if err != nil {
			return err
		}

		artistRepo := NewArtistRepository(tx)
		err = artistRepo.DeleteArtistMappingWithMovieId(ctx, movie.ID)
		if err != nil {
			return err
		}
		for _, artist := range artists {
			err = artistRepo.UpsertArtist(ctx, artist)
			if err != nil {
				return err
			}

			err = artistRepo.InsertArtistWithMovieMapping(ctx, artist, movie.ID)
			if err != nil {
				return err
			}
		}

		genreRepo := NewGenreRepository(tx)
		err = genreRepo.DeleteGenreMappingWithMovieId(ctx, movie.ID)
		if err != nil {
			return err
		}
		for _, genre := range genres {
			err = genreRepo.UpsertGenre(ctx, genre)
			if err != nil {
				return err
			}

			err = genreRepo.InsertGenreWithMovieMapping(ctx, genre, movie.ID)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (ur *movieRepository) CreateAndScanId(ctx context.Context, movie *model.Movie) error {

	sql := `INSERT INTO movies
		(title, description, watch_url, duration)
		VALUES( $1, $2, $3, $4) RETURNING id;`

	return ur.db.
		QueryRow(ctx, sql, movie.Title, movie.Description, movie.WatchUrl, movie.Duration).
		Scan(&movie.ID)
}

func (ur *movieRepository) GetMovieDetail(ctx context.Context, movieID int) (*model.Movie, error) {

	sql := `SELECT 
				m.id,
				m.title,
				m.description,
				m.watch_url,
				m.duration,
				coalesce (mv."views", 0) as views,
				JSON_AGG(distinct ma.artist_name) as artists,
				JSON_AGG(distinct mg.genre_name) as genres
			FROM 
				movies m 
			left join 
				movie_artists ma on m.id = ma.movie_id 
			left join 
				movie_genres mg on m.id = mg.movie_id 
			left join
				movie_views mv on m.id = mv.movie_id
			where m.id = $1 group by m.id, mv."views" limit 1`

	genres := []string{}
	artists := []string{}
	movie := &model.Movie{
		Genres: []string{},
		Artist: []string{},
	}
	err := ur.db.QueryRow(ctx, sql, movieID).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.WatchUrl,
		&movie.Duration,
		&movie.Views,
		&artists,
		&genres,
	)
	if err != nil {
		return nil, err
	}

	for _, genre := range genres {
		if genre != "" {
			movie.Genres = append(movie.Genres, genre)
		}
	}

	for _, artist := range artists {
		if artist != "" {
			movie.Artist = append(movie.Artist, artist)
		}
	}

	return movie, nil
}

func (ur *movieRepository) GetMoviesWithPaginationAndSearch(ctx context.Context, search, searchBy, orderBy, orderDirection string, page, pageSize int) ([]model.Movie, error) {
	q := `SELECT 
			m.id,
			m.title,
			m.description,
			m.watch_url,
			m.duration,
			coalesce (mv."views", 0) as views,
			JSON_AGG(distinct ma.artist_name) as artists,
			JSON_AGG(distinct mg.genre_name) as genres
		FROM 
			movies m 
		left join 
			movie_artists ma on m.id = ma.movie_id 
		left join 
			movie_genres mg on m.id = mg.movie_id 
		left join
			movie_views mv on m.id = mv.movie_id`

	if search != "" {
		switch search {
		case "title":
			q = q + "where m.title = $1"
		case "description":
			q = q + "where m.description = $1"
		case "artists":
			q = q + "where ma.artist_name = $1"
		case "genres":
			q = q + "where ma.genre_name = $1"
		default:
			return nil, errors.New("search can be processed by title/description/artists/genres")
		}
	}

	q = q + ` group by m.id, mv."views"`

	if orderBy != "" {
		switch orderDirection {
		case "asc":
			orderDirection = " asc"
		default:
			orderDirection = " desc"
		}

		switch orderBy {
		case "title":
			q = q + ` order by "m.title"` + orderDirection
		case "description":
			q = q + ` order by "m.description"` + orderDirection
		case "views":
			q = q + ` order by "views"` + orderDirection
		default:
			return nil, errors.New("search can be processed by title/description/artists/genres")
		}
	}

	if pageSize == 0 {
		pageSize = 10
	}

	q = q + fmt.Sprintf(` limit %d`, pageSize)
	if page > 0 {
		skip := (page - 1) * pageSize
		q = q + fmt.Sprintf(` offset %d`, skip)
	}

	rows, err := ur.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	movies := []model.Movie{}
	defer rows.Close()
	for rows.Next() {
		genres := []string{}
		artists := []string{}
		movie := &model.Movie{
			Genres: []string{},
			Artist: []string{},
		}
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.WatchUrl,
			&movie.Duration,
			&movie.Views,
			&artists,
			&genres,
		)
		if err != nil {
			return nil, err
		}

		for _, genre := range genres {
			if genre != "" {
				movie.Genres = append(movie.Genres, genre)
			}
		}

		for _, artist := range artists {
			if artist != "" {
				movie.Artist = append(movie.Artist, artist)
			}
		}

		movies = append(movies, *movie)
	}

	return movies, nil
}
