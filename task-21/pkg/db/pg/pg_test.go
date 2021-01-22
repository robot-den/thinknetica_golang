package pg

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"reflect"
	"sort"
	"task-21/pkg/model"
	"testing"
)

var pg *PG
var dbpool *pgxpool.Pool
var ctx = context.TODO()

func TestMain(m *testing.M) {
	url := "postgresql://golang:golang@localhost:54321/golang"
	dbpool, _ = pgxpool.Connect(ctx, url)
	pg = New(dbpool)

	exitCode := m.Run()
	dbpool.Close()

	os.Exit(exitCode)
}

func TestPG_InsertMovies(t *testing.T) {
	movies := []model.Movie{
		{Name: "Some movie", ReleaseYear: 1990, Rating: "PG-18", Gross: 5000000, StudioID: 1},
	}
	err := pg.InsertMovies(ctx, movies)
	if err != nil {
		t.Errorf("pg.InsertMovies(); err = %v, want %v", err, nil)
	}
}

func TestPG_DeleteMovie(t *testing.T) {
	var id int
	_ = dbpool.QueryRow(ctx, "SELECT id FROM movies WHERE name = 'Some movie' LIMIT 1").Scan(&id)

	err := pg.DeleteMovie(ctx, id)
	if err != nil {
		t.Errorf("pg.DeleteMovie(); err = %v, want %v", err, nil)
	}
}

func TestPG_UpdateMovie(t *testing.T) {
	// Select movie before update to return state back
	movie := model.Movie{}
	row := dbpool.QueryRow(ctx, "SELECT * FROM movies WHERE id = 1 LIMIT 1")
	_ = row.Scan(&movie.ID, &movie.Name, &movie.ReleaseYear, &movie.Rating, &movie.Gross, &movie.StudioID)

	updatedMovie := model.Movie{
		Name: "Some movie", ReleaseYear: 1990, Rating: "PG-18", Gross: 5000000, StudioID: 1, ID: movie.ID,
	}

	err := pg.UpdateMovie(ctx, updatedMovie)
	if err != nil {
		t.Errorf("pg.UpdateMovie(); err = %v, want %v", err, nil)
	}

	_ = pg.UpdateMovie(ctx, movie)
}

func TestPG_SelectMovies(t *testing.T) {
	movies, err := pg.SelectMovies(ctx, 0)
	if err != nil {
		t.Errorf("pg.SelectMovies(); err = %v, want %v", err, nil)
	}

	var ids []int
	for _, movie := range movies {
		ids = append(ids, movie.ID)
	}
	sort.Ints(ids)

	got := ids
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("!reflect.DeepEqual(got, want) = %v; want %v", got, want)
	}
}

func TestPG_SelectMovies_by_studio(t *testing.T) {
	movies, err := pg.SelectMovies(ctx, 2)
	if err != nil {
		t.Errorf("pg.SelectMovies(); err = %v, want %v", err, nil)
	}

	var ids []int
	for _, movie := range movies {
		ids = append(ids, movie.ID)
	}
	sort.Ints(ids)

	got := ids
	want := []int{3, 4, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("!reflect.DeepEqual(got, want) = %v; want %v", got, want)
	}
}
