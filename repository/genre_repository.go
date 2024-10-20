package repository

import (
	"context"

	"github.com/elangreza14/moviefestival/model"
)

type genreRepository struct {
	db QueryPgx
	*PostgresRepo[model.Genre]
}

func NewGenreRepository(
	dbPool QueryPgx,
) *genreRepository {
	return &genreRepository{
		db:           dbPool,
		PostgresRepo: NewPostgresRepo[model.Genre](dbPool),
	}
}

func (ar *genreRepository) UpsertGenre(ctx context.Context, name string) error {
	sql := `INSERT INTO genres ("name") VALUES ($1) ON CONFLICT (name) DO NOTHING`
	_, err := ar.db.Exec(ctx, sql, name)
	if err != nil {
		return err
	}
	return err
}

func (ar *genreRepository) InsertGenreWithMovieMapping(ctx context.Context, name string, movieID int) error {
	sql := `INSERT INTO movie_genres ("movie_id", "genre_name") VALUES ($1, $2)
			ON CONFLICT (movie_id,genre_name) DO NOTHING`
	_, err := ar.db.Exec(ctx, sql, movieID, name)
	if err != nil {
		return err
	}
	return err
}

func (ar *genreRepository) DeleteGenreMappingWithMovieId(ctx context.Context, movieID int) error {
	sql := `DELETE FROM movie_genres WHERE movie_id=$1`
	_, err := ar.db.Exec(ctx, sql, movieID)
	if err != nil {
		return err
	}
	return err
}

type ViewedGenre struct {
	Name  string `db:"name"`
	Views int    `db:"views"`
}

func (ur *genreRepository) GetMostViewedGenres(ctx context.Context) ([]model.ViewedGenre, error) {
	sql := `select 
				mg.genre_name, SUM(mv."views") as views 
			from 
				movie_views mv join movie_genres mg 
			on 
				mv.movie_id = mg.movie_id 
			group by 
				mg.genre_name 
			order by 
				views desc limit 10`

	rows, err := ur.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	res := []model.ViewedGenre{}
	for rows.Next() {
		genre := model.ViewedGenre{}
		err = rows.Scan(&genre.Name, &genre.Views)
		if err != nil {
			return nil, err
		}

		res = append(res, genre)
	}

	return res, nil
}
