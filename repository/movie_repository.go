package repository

import (
	"context"

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
			join 
				movie_artists ma on m.id = ma.movie_id 
			join 
				movie_genres mg on m.id = mg.movie_id 
			left join
				movie_views mv on m.id = mv.movie_id
			where m.id = $1 group by m.id, mv."views" limit 1`

	movie := &model.Movie{}
	err := ur.db.QueryRow(ctx, sql, movieID).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.WatchUrl,
		&movie.Duration,
		&movie.Views,
		&movie.Artist,
		&movie.Genres)
	if err != nil {
		return nil, err
	}

	return movie, err
}
