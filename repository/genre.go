package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/model"
)

func AllGenre(db *sqlx.DB) ([]model.Genre, error) {
	g := make([]model.Genre, 0)
	if err := db.Select(&g, `SELECT id, genre_abbreviation_name, genre_name FROM genre`); err != nil {
		return nil, err
	}

	return g, nil
}

func WriteGenre(db *sqlx.Tx, g *model.Genre) (result sql.Result, err error) {
	stmt, err := db.Prepare(`INSERT INTO genre (genre_abbreviation_name, genre_name) VALUES (?, ?)`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	return stmt.Exec(g.GenreAbbrName, g.GenreName)
}