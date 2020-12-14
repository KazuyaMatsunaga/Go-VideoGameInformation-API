package model

import (
	"time"
)

type Game struct {
	ID int64 `db:"id"`
	Title string `db:"title" json:"title"`
	Price string `db:"price" json:"price"`
	ReleaseDate time.Time `db:"release_date" json:"release_date"`
}