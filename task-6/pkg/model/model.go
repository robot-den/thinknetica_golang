// Package model предоставляет структуры данных которые используются в нескольких пакетах приложения
package model

type Record struct {
	Id    int
	Url   string
	Title string
}

type InvertedIndex map[string][]int
