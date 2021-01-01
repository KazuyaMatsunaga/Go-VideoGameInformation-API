package scrapingtask

import (
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/repository"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/service"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/model"
)

func scrapingPf() []model.Platform {
	var pfList []model.Platform
	repo := repository.NewPlatformClient()
	s := service.NewPlatformService(repo)
	pfList = s.Platform()

	return pfList
}