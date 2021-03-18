package controller

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/repository"
)

type Platform struct {
	db *sqlx.DB
}

func NewPlatform(db *sqlx.DB) *Platform {
	return &Platform{db}
}

func (p *Platform) Index(_ http.ResponseWriter, _ *http.Request) (int, interface{}, error) {
	Platforms, err := repository.AllPf(p.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, Platforms, nil
}
