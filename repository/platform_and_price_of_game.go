package repository

import (
	"database/sql"

	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/model"
	"github.com/jmoiron/sqlx"
)

func WritePfAndPrice(db *sqlx.Tx, ga *model.Game, pf *model.Platform, price string) (result sql.Result, err error) {
	stmt, err := db.Prepare(`INSERT INTO platform_and_price_of_game (game_id, platform_id, price) VALUES (?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	return stmt.Exec(ga.ID, pf.ID, price)
}
