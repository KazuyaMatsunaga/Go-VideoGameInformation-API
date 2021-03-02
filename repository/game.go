package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/model"
)

func AllGame(db *sqlx.DB) ([]model.Game, error) {
	ga := make([]model.Game, 0)
	if err := db.Select(&ga, `SELECT id, title FROM game`); err != nil {
		return nil, err
	}

	return ga, nil
}

func GetGame(db *sqlx.DB, ga *model.Game) (*model.Game, error) {
	var gav model.Game
	if err := db.Get(&gav,`SELECT id, title FROM game WHERE id = ? LIMIT 1`, ga.ID); err != nil {
		return nil, err
	}
	return &gav, nil
}

func WriteGame(db *sqlx.Tx, ga *model.Game) (result sql.Result, err error) {
	stmt, err := db.Prepare(`INSERT INTO game (title) VALUES (?)`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	return stmt.Exec(ga.Title)
}