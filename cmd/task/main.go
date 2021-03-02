package main

import (
	"fmt"
	"flag"

	spRepo "github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/repository"
	spSVC "github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/service"

	PkgRepo "github.com/KazuyaMatsunaga/Go-VideoGame-Package-Search/pkg/repository"
	PkgSvc "github.com/KazuyaMatsunaga/Go-VideoGame-Package-Search/pkg/service"

	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/model"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/service"
	"github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/db"
)

var (
	info = flag.String("info", "", "target info for scraping")
)

func main() {
	var databaseDatasource string
	flag.StringVar(&databaseDatasource, "databaseDatasource", "root:password@tcp(db:3306)/game_information", "Should looks like root:password@tcp(hostname:port)/dbname")
	flag.Parse()
	cs := db.NewDB(databaseDatasource)
	dbcon, err := cs.Open()
	if err != nil {
		fmt.Errorf("failed db init. %s", err)
	}

	switch *info {
	case "genre":
		repo := spRepo.NewGenreClient()
		s := spSVC.NewGenreService(repo)
		genreList := s.Genre()
		for _, g := range genreList {
			newGenre := model.Genre{
				GenreAbbrName: g.Addr,
				GenreName: g.Name,
			}
			if newGenre.GenreAbbrName == "" || newGenre.GenreName == "" {
				continue
			}
			genreService := service.NewGenre(dbcon)
			_, err := genreService.Write(&newGenre)
			if err != nil {
				fmt.Errorf("failed write for genre. %s\n", err)
			}
		}
		fmt.Println("finish write for genre.\n")
	case "platform":
		repo := spRepo.NewPlatformClient()
		s := spSVC.NewPlatformService(repo)
		pfList := s.Platform()
		for _, p := range pfList {
			newPf := model.Platform{
				PfAbbrName: p.Addr,
				PfName: p.Name,
			}
			if newPf.PfAbbrName == "" || newPf.PfName == "" {
				continue
			}
			pfService := service.NewPlatform(dbcon)
			_, err := pfService.Write(&newPf)
			if err != nil {
				fmt.Errorf("failed write for platform. %s\n", err)
			}
		}
		fmt.Println("finish write for platform.\n")
	case "detail":
		repo := spRepo.NewDetailClient()
		s := spSVC.NewDetailService(repo)
		details := s.Detail()
		repoPkg := PkgRepo.NewPackageClient()
		sP := PkgSvc.NewPackageService(repoPkg)
		detailList := sP.Package(details)
		fmt.Printf("%v\n", detailList)
	}
}