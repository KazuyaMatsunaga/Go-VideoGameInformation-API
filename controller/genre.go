package controller

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/repository"
)

type Genre struct {
	db *sqlx.DB
}

func NewGenre(db *sqlx.DB) *Genre {
	return &Genre{db}
}

func (g *Genre) Index(_ http.ResponseWriter, _ *http.Request) (int, interface{}, error) {
	genres, err := repository.AllGenre(g.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, genres, nil
}
