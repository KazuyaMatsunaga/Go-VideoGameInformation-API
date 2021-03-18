package model

type Genre struct {
	ID            int64  `db:"id" json:"id"`
	GenreAbbrName string `db:"genre_abbreviation_name" json:"genre_abbreviation_name"`
	GenreName     string `db:"genre_name" json:"genre_name"`
}
