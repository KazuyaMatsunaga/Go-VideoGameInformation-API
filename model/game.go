package model

import (
	"time"
)

type Game struct {
	ID int64 `db:"id"`
	Price int64 `db:"price" json:"price"`
	ReleaseDate time.Time `db:"release_date" json:"release_date"`
	TitleID int64 `db:"title_id" json:"title_id"`
	PlatformID int64 `db:"platform_id" json:"platform_id"`
	GenreID int64 `db:"genre_id" json:"genre_id"`
	PackageImgID int64 `db:"package_image_id" json:"package_image_id"`
	ExecutedScrapinigID int64 `db:"executed_scraping_id" json:"executed_scraping_id"`
}