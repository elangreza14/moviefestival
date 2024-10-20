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
