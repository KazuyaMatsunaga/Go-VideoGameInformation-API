package scrapingtask

import (
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/repository"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/service"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/model"
)

func scrapingGenre() []model.Genre {
	var genreList []model.Genre
	repo := repository.NewGenreClient()
	s := service.NewGenreService(repo)
	genreList = s.Genre()

	return genreList
}