package scrapingutil

import (
	"log"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/Go-VideoGameInformation-API/scrapingmodel"
)

var gameCatURL = "https://w.atwiki.jp/gcmatome/"

func OpenURL(url string) (*goquery.Document, error) {
	doc, openErr := goquery.NewDocument(url)
	if openErr != nil {
		log.Print(openErr)
	}
	return doc, openErr
}

func ExtraReplacer(str string) string {
	var afterStr string

	afterStr = strings.Replace(str," ","",-1)
	afterStr = strings.Replace(str,"\n","",-1)

	return afterStr
}

func FindGamePlatformURL(gamePf string) string {
	doc, _ := OpenURL(gameCatURL)

	var pfURL string

	linkTitle := "ゲーム記事一覧 (" + gamePf + ")"

	doc.Find("ul.pl-l-level-2 a[title]").Each(func(_ int, s *goquery.Selection) {
		titleAttr, _ := s.Attr("title")
		if c := strings.Contains(titleAttr, linkTitle); c {
			hrefStr, _ := s.Attr("href")
			pfURL = "https:" + hrefStr
		}
	})

	return pfURL
}

func FindGameList(pfURL string, searchQry string) []scrapingmodel.Game {
	docGameList, _ := OpenURL(pfURL)

	g := scrapingmodel.Game{}

	gS := make([]scrapingmodel.Game,0)

	docGameList.Find("table[cellspacing] a[title]").Each(func(_ int, s *goquery.Selection) {
		if  searchQry != "" {
			titleAttr, _ := s.Attr("title")
			if c := strings.Contains(titleAttr, searchQry); c {
				genreStr := s.ParentFiltered("td").Next().Text()
				genreStr = ExtraReplacer(genreStr)
				titleStr := s.Text()
				titleStr = ExtraReplacer(titleStr)
				hrefAttr , _ := s.Attr("href")
				hrefStr := "https:" + hrefAttr
				g.Title = titleStr
				g.URL = hrefStr
				g.Genre = genreStr
			}
		} else {
			genreStr := s.ParentFiltered("td").Next().Text()
			genreStr = ExtraReplacer(genreStr)
			titleStr := s.Text()
			titleStr = ExtraReplacer(titleStr)
			hrefAttr , _ := s.Attr("href")
			hrefStr := "https:" + hrefAttr
			g.Title = titleStr
			g.URL = hrefStr
			g.Genre = genreStr
		}
		gS = append(gS, g)
	})

	return gS
}

func SetGameDetail(gS []scrapingmodel.Game) []scrapingmodel.Game {
	for i, g := range gS {
		docGame, _ := OpenURL(g.URL)


	}
}



