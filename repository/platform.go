package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/model"
)

func AllPf(db *sqlx.DB) ([]model.Platform, error) {
	p := make([]model.Platform, 0)
	if err := db.Select(&p, `SELECT id, platform_abbreviation_name, platform_name FROM platform`); err != nil {
		return nil, err
	}

	return p, nil
}

func WritePf(db *sqlx.Tx, p *model.Platform) (result sql.Result, err error) {
	stmt, err := db.Prepare(`INSERT INTO platform (platform_abbreviation_name, platform_name) VALUES (?, ?)`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	return stmt.Exec(p.PfAbbrName, p.PfName)
}