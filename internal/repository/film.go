package repository

import (
	"context"
	"errors"
	"film-library/internal/db"
	"film-library/internal/model"
)

type FilmRepo struct {
	db *db.Database
}

func NewFilms(database *db.Database) *FilmRepo {
	return &FilmRepo{db: database}
}

func (f *FilmRepo) AddFilm(ctx context.Context, film model.Film) (int64, error) {
	var id int64
	err := f.db.ExecQueryRow(ctx, `INSERT INTO films (name,description,rating,releasedate) VALUES ($1,$2,$3,$4) returning id`, film.Name, film.Description, film.Rating, film.ReleaseDate).Scan(&id)
	if err != nil {
		return -1, errors.New("can not exec insert")
	}
	return id, nil
}

func (f *FilmRepo) DeleteFilm(ctx context.Context, film model.Film) (int64, error) {
	_, err := f.db.Exec(ctx, `DELETE FROM films WHERE id = $1`, film.ID)
	if err != nil {
		return -1, errors.New("can not exec delete")
	}
	return film.ID, nil
}

func (f *FilmRepo) SelectFilm(ctx context.Context, film model.Film) (*model.Film, error) {
	r := f.db.ExecQueryRow(ctx, `SELECT * FROM films WHERE id = $1`, film.ID)
	convfilm := model.Film{}
	err := r.Scan(&convfilm)
	if err != nil {
		return &model.Film{}, errors.New("can not convert to this structure")
	}
	return &convfilm, nil
}

func (f *FilmRepo) SortFilms(ctx context.Context, sortparam string) (*[]model.Film, error) {
	r := f.db.ExecQueryRow(ctx, `SELECT * FROM films ORDER BY $1 DESC`, sortparam)
	convfilms := []model.Film{}
	err := r.Scan(&convfilms)
	if err != nil {
		return nil, errors.New("can not convert to this structure")
	}
	return &convfilms, nil
}

func (f *FilmRepo) UpdateFilm(ctx context.Context, film model.Film) error {
	_, err := f.db.Exec(ctx, `UPDATE films SET name = $2, description = $3, rating = $4, releasedate = $5 WHERE id = $1`, film.ID, film.Name, film.Description, film.Rating, film.ReleaseDate)
	if err != nil {
		return errors.New("can not convert to this structure")
	}
	return nil
}
