package model

type GenreOfGame struct{
	GameID int64 `db:"game_id"`
	GenreID int64 `db:"genre_id"`
}