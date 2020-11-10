package model

import (
	"time"
)

type Game struct {
	ID int64 `db:"id"`
	Title string `db:"title" json:"title"`
	Price int64 `db:"price" json:"price"`
	ReleaseDate time.Time `db:"release_date" json:"release_date"`
	WikiURL string `db:"wiki_url" json:"wiki_url"`
	PlatformID int64 `db:"platform_id" json:"platform_id"`
	GenreID int64 `db:"genre_id" json:"genre_id"`
	PackageImageID int64 `db:"package_image_id" json:"package_image_id"`
	ExecutedScrapinigID int64 `db:"executed_scraping_id" json:"executed_scraping_id"`
}