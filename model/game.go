package model

type Game struct {
	ID int64 `db:"id"`
	Title string `db:"title" json:"title"`
}