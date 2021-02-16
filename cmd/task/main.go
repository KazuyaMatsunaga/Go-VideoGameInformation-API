package main

import (
	"fmt"
	"flag"

	spRepo "github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/repository"
	spSVC "github.com/KazuyaMatsunaga/Go-VideoGameInformation-Scraping/pkg/service"

	PkgRepo "github.com/KazuyaMatsunaga/Go-VideoGame-Package-Search/pkg/repository"
	PkgSvc "github.com/KazuyaMatsunaga/Go-VideoGame-Package-Search/pkg/service"
)

var (
	info = flag.String("info", "", "target info for scraping")
)

func main() {
	flag.Parse()

	switch *info {
	case "genre":
		repo := spRepo.NewGenreClient()
		s := spSVC.NewGenreService(repo)
		genreList := s.Genre()
		fmt.Printf("%v\n", genreList)
	case "platform":
		repo := spRepo.NewPlatformClient()
		s := spSVC.NewPlatformService(repo)
		pfList := s.Platform()
		fmt.Printf("%v\n", pfList)
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