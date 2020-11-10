package model

type Genre struct{
	ID int64 `db:"id"`
	GenreName string `db:"genre_name" json:"genre_name"`
}