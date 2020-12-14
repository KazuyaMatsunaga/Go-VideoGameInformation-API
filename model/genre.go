package model

import (
	"time"
)

type Genre struct{
	ID int64 `db:"id"`
	GenreAbbrName string `db:"genre_abbreviation_name" json:"genre_abbreviation_name"`
	GenreName string `db:"genre_name" json:"genre_name"`
	CreatedTime      time.Time   `db:"ctime" json:"created_time"`
	UpdatedTime      time.Time   `db:"utime" json:"updated_time"`
}