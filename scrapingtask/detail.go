package scrapingtask

import (
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/repository"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/service"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/model"
)

func scrapingDetail() []model.Detail {
	var detailList []model.Detail
	repo := repository.NewDetailClient()
	s := service.NewDetailService(repo)
	detailList = s.Detail()

	return detailList
}