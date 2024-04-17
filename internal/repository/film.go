package repository

import (
	"context"
	"database/sql"
	"errors"
	"film-library/internal/db"
	"film-library/internal/model"
	"time"
)

type FilmRepo struct {
	db db.DBops
}

func NewRepository(database db.DBops) *FilmRepo {
	return &FilmRepo{db: database}
}

func (f *FilmRepo) AddFilm(ctx context.Context, film model.Film) (int64, error) {
	var id int64
	err := f.db.ExecQueryRow(ctx, `INSERT INTO films (name,description,rating,releasedate) VALUES ($1,$2,$3,$4) returning id`, film.Name, film.Description, film.Rating, parseDate(film.ReleaseDate)).Scan(&id)
	if err != nil {
		return -1, errors.New("can not exec insert")
	}
	return id, nil
}

func (f *FilmRepo) DeleteFilm(ctx context.Context, queryid int64) (int64, error) {
	_, err := f.db.Exec(ctx, `DELETE FROM films WHERE id = $1`, queryid)
	if err != nil {
		return -1, errors.New("can not exec delete")
	}
	return queryid, nil
}

func (f *FilmRepo) SelectFilm(ctx context.Context, queryid int64) (*model.Film, error) {
	var convfilm model.Film
	err := f.db.Get(ctx, &convfilm, `SELECT id,name,description,rating,releasedate FROM films WHERE id = $1`, queryid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrObjectNotFound
		}
		return nil, err
	}
	return &convfilm, nil
}

func (f *FilmRepo) SortFilms(ctx context.Context, sortparam string) ([]model.Film, error) {
	convfilms := []model.Film{}
	switch sortparam {
	case "name":
		err := f.db.Select(ctx, &convfilms, `SELECT name,description,rating,releasedate FROM films ORDER BY name DESC`)
		if err != nil {
			return nil, errors.New("can not convert to this structure")
		}
	case "description":
		err := f.db.Select(ctx, &convfilms, `SELECT name,description,rating,releasedate FROM films ORDER BY description DESC`)
		if err != nil {
			return nil, errors.New("can not convert to this structure")
		}
	case "rating":
		err := f.db.Select(ctx, &convfilms, `SELECT name,description,rating,releasedate FROM films ORDER BY rating DESC`)
		if err != nil {
			return nil, errors.New("can not convert to this structure")
		}
	case "releasedate":
		err := f.db.Select(ctx, &convfilms, `SELECT name,description,rating,releasedate FROM films ORDER BY releasedate DESC`)
		if err != nil {
			return nil, errors.New("can not convert to this structure")
		}
	default:
		return nil, errors.New("wrong param")
	}
	return convfilms, nil
}

//func (f *FilmRepo) UpdateFilm(ctx context.Context, film model.Film) error {
//	_, err := f.db.Exec(ctx, `UPDATE films SET name = $2, description = $3, rating = $4, releasedate = $5 WHERE id = $1`, 50000, film.Name, film.Description, film.Rating, film.ReleaseDate) //TODO!!!!
//	if err != nil {
//		return errors.New("can not convert to this structure")
//	}
//	return nil
//}

func parseDate(t time.Time) string {
	return t.Format("2006-01-02")
}
