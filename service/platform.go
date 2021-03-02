package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/dbutil"

	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/model"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/repository"
)

type Platform struct {
	db *sqlx.DB
}

func NewPlatform(db *sqlx.DB) *Platform {
	return &Platform{db}
}

func (p *Platform) Write(pf *model.Platform) (int64, error) {
	var createdID int64
	if err := dbutil.TXHandler(p.db, func(tx *sqlx.Tx) error {
		result, err := repository.WritePf(tx, pf)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		createdID = id
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed Platform insert transaction")
	}
	return createdID, nil
}

