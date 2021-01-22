// Package db содержит интерфейс базы данных
package db

import (
	"context"
	"task-21/pkg/model"
)

// DB определяет контракт базы данных
type DB interface {
	InsertMovies(context.Context, []model.Movie) error
	DeleteMovie(context.Context, int) error
	UpdateMovie(context.Context, model.Movie) error
	SelectMovies(context.Context, int) ([]model.Movie, error)
}
