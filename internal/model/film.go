package model

import "time"

type Film struct {
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Rating      int64     `json:"rating" db:"rating"`
	ReleaseDate time.Time `json:"releasedate" db:"releasedate"`
}

type Actor struct {
	Name      string    `db:"actorname"`
	Gender    string    `db:"gender"`
	BirthDate time.Time `db:"birthdate"`
}
