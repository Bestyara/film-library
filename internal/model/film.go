package model

import "time"

type Film struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Rating      int64     `json:"rating" db:"rating"`
	ReleaseDate time.Time `json:"releasedate" db:"releasedate"`
}
