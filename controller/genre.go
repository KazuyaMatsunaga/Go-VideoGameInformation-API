package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/model"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/repository"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/service"
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

func (g *Genre) Write(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newGenre := &model.Genre{}
	if err := json.NewDecoder(r.Body).Decode(&newGenre); err != nil {
		return http.StatusBadRequest, nil, err
	}

	if newGenre.GenreAbbrName == "" || newGenre.GenreName == "" {
		return http.StatusUnprocessableEntity, nil, errors.New("required parameter is missing")
	}

	genreService := service.NewGenre(g.db)
	id, err := genreService.Write(newGenre)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newGenre.ID = id
	return http.StatusCreated, newGenre, nil
}
