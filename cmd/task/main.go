package main

import (
	"fmt"
	"flag"

	scrapingtask "github.com/KazuyaMatsunaga/Go-VideoGameInformation-API/scrapingtask"
)

var (
	info = flag.String("info", "", "target info for scraping")
)

func main() {
	flag.Parse()

	switch *info {
	case "genre":
		genreList := scrapingtask.scrapingGenre()
		fmt.Printf("%v\n", genreList)
	case "platform":
		pfList := scrapingtask.scrapingPf()
		fmt.Printf("%v\n", pfList)
	case "detail":
		detailList := scrapingtask.scrapingDetail()
		fmt.Printf("%v\n", detailList)
	}
}