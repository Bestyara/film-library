package service

import (
	"context"
	"film-library/internal/model"
	"log"
)

type repos interface {
	AddFilm(ctx context.Context, film model.Film) (int64, error)
	DeleteFilm(ctx context.Context, queryid int64) (int64, error)
	SelectFilm(ctx context.Context, film model.Film) (*model.Film, error)
	SortFilms(ctx context.Context, sortparam string) ([]model.Film, error)
	UpdateFilm(ctx context.Context, film model.Film) error
}

type FilmServ struct {
	repo repos
}

func NewService(r repos) *FilmServ {
	return &FilmServ{repo: r}
}

func (f *FilmServ) AddFilm(ctx context.Context, film model.Film) (int64, error) {
	if film.Rating < 0 && film.Rating > 10 {
		log.Println("Wrong input of rating")
	}
	if len([]rune(film.Name)) <= 0 && len([]rune(film.Name)) > 150 {
		log.Println("Wrong length of name")
	}
	if len([]rune(film.Description)) > 1000 {
		log.Println("Wrong length of description")
	}
	return f.repo.AddFilm(ctx, film)
}

func (f *FilmServ) DeleteFilm(ctx context.Context, queryid int64) (int64, error) {
	if queryid < 0 {
		log.Println("Negative queryid")
	}
	return f.repo.DeleteFilm(ctx, queryid)
}
func (f *FilmServ) SelectFilm(ctx context.Context, film model.Film) (*model.Film, error) {
	return f.repo.SelectFilm(ctx, film)
}
func (f *FilmServ) SortFilms(ctx context.Context, sortparam string) ([]model.Film, error) {
	if sortparam == "default" {
		return f.repo.SortFilms(ctx, "rating") //если параметр сортировки не передан, то сортируем по рейтингу
	}
	return f.repo.SortFilms(ctx, sortparam)
}
func (f *FilmServ) UpdateFilm(ctx context.Context, film model.Film) error {
	if film.Rating < 0 && film.Rating > 10 {
		log.Println("Wrong input of rating")
	}
	if len([]rune(film.Name)) <= 0 && len([]rune(film.Name)) > 150 {
		log.Println("Wrong length of name")
	}
	if len([]rune(film.Description)) > 1000 {
		log.Println("Wrong length of description")
	}
	return f.repo.UpdateFilm(ctx, film)
}
