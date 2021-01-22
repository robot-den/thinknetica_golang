// package pg представляет собой сервис для работы с базой фильмов
package pg

import (
	"context"
	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"task-21/pkg/model"
)

// PG представляет собой тип, реализующий функции по работе с базой фильмов
type PG struct {
	pool *pgxpool.Pool
}

// New создает новый объект PG
func New(pool *pgxpool.Pool) *PG {
	p := PG{
		pool: pool,
	}
	return &p
}

// InsertMovies добавляет фильмы в базу с использованием транзакции
func (p *PG) InsertMovies(ctx context.Context, movies []model.Movie) error {
	sql := "INSERT INTO movies (name, release_year, rating, gross, studio_id) VALUES ($1, $2, $3, $4, $5)"
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	batch := new(pgx.Batch)
	for _, movie := range movies {
		batch.Queue(sql, movie.Name, movie.ReleaseYear, movie.Rating, movie.Gross, movie.StudioID)
	}

	results := tx.SendBatch(ctx, batch)
	err = results.Close()
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// DeleteMovie выполняет удаление фильма из базы
func (p *PG) DeleteMovie(ctx context.Context, id int) error {
	_, err := p.pool.Exec(ctx, "DELETE FROM movies WHERE id = $1", id)
	return err
}

// UpdateMovie обновляет фильм в базе
func (p *PG) UpdateMovie(ctx context.Context, movie model.Movie) error {
	sql := "UPDATE movies SET name = $1, release_year = $2, rating = $3, gross = $4, studio_id = $5 WHERE id = $6"
	_, err := p.pool.Exec(ctx, sql, movie.Name, movie.ReleaseYear, movie.Rating, movie.Gross, movie.StudioID, movie.ID)
	return err
}

// SelectMovies выбирает из базы фильмы указанной студии (если StudioID равен нулю, то будут выбраны все фильмы)
func (p *PG) SelectMovies(ctx context.Context, StudioID int) ([]model.Movie, error) {
	var movies []model.Movie
	var err error
	var rows pgx.Rows

	rows, err = p.pool.Query(ctx, "SELECT * from MOVIES WHERE studio_id = $1 OR $1 = 0", StudioID)

	if err != nil {
		return movies, err
	}

	for rows.Next() {
		var m model.Movie
		err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.ReleaseYear,
			&m.Rating,
			&m.Gross,
			&m.StudioID,
		)
		if err != nil {
			return movies, err
		}
		movies = append(movies, m)
	}

	err = rows.Err()
	if err != nil {
		return movies, err
	}

	return movies, nil
}
