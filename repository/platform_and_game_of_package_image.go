package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/model"
)

func WritePfAndPackageImg(db *sqlx.Tx, ga *model.Game, pf *model.Platform, pkg *model.PackageImg) (result sql.Result, err error) {
	stmt, err := db.Prepare(`INSERT INTO platform_and_package_image_of_game (game_id, platform_id, package_image_id) VALUES (?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	return stmt.Exec(ga.ID, pf.ID, pkg.ID)
}