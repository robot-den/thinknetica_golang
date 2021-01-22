// package model используется для хранения типов, задействованных в разных частях приложения
package model

// Movie представляет собой фильм
type Movie struct {
	ID          int
	Name        string
	ReleaseYear int
	Rating      string
	Gross       int
	StudioID    int
}
