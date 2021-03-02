package repository

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/model"
)

func WritePfAndRlDate(db *sqlx.Tx, ga *model.Game, pf *model.Platform, rlDate time.Time) (result sql.Result, err error) {
	stmt, err := db.Prepare(`INSERT INTO platform_and_release_date_of_game (game_id, platform_id, release_date) VALUES (?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	return stmt.Exec(ga.ID, pf.ID, rlDate)
}