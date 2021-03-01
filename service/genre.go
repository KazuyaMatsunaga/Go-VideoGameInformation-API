package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/dbutil"

	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/model"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/repository"
)

type Genre struct {
	db *sqlx.DB
}

func NewGenre(db *sqlx.DB) *Genre {
	return &Genre{db}
}

func (g *Genre) Write(genre *model.Genre) (int64, error) {
	var createdID int64
	if err := dbutil.TXHandler(g.db, func(tx *sqlx.Tx) error {
		result, err := repository.WriteGenre(tx, genre)
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
		return 0, errors.Wrap(err, "failed genre insert transaction")
	}
	return createdID, nil
}

