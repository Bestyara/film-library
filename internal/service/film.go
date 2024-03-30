package service

import (
	"context"
	"film-library/internal/model"
)

type repos interface {
	AddFilm(ctx context.Context, film model.Film) (int64, error)
	DeleteFilm(ctx context.Context, film model.Film) (int64, error)
	SelectFilm(ctx context.Context, film model.Film) (*model.Film, error)
	SortFilms(ctx context.Context, sortparam string) (*[]model.Film, error)
	UpdateFilm(ctx context.Context, film model.Film) error
}

type FilmServ struct {
	repo repos
}

func NewService(r repos) *FilmServ {
	return &FilmServ{repo: r}
}

func (f *FilmServ) AddFilm(ctx context.Context, film model.Film) (int64, error) {
	return f.repo.AddFilm(ctx, film)
}

func (f *FilmServ) DeleteFilm(ctx context.Context, film model.Film) (int64, error) {
	return f.repo.DeleteFilm(ctx, film)
}
func (f *FilmServ) SelectFilm(ctx context.Context, film model.Film) (*model.Film, error) {
	return f.repo.SelectFilm(ctx, film)
}
func (f *FilmServ) SortFilms(ctx context.Context, sortparam string) (*[]model.Film, error) {
	return f.repo.SortFilms(ctx, sortparam)
}
func (f *FilmServ) UpdateFilm(ctx context.Context, film model.Film) error {
	return f.repo.UpdateFilm(ctx, film)
}
