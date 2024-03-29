package model

import "time"

type Film struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Rating      int64     `db:"rating"`
	ReleaseDate time.Time `db:"releasedate"`
}

type Actor struct {
	Name      string    `db:"actorname"`
	Gender    string    `db:"gender"`
	BirthDate time.Time `db:"birthdate"`
}
