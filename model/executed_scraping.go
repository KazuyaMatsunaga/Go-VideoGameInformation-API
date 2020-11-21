package model

import (
	"time"
)

type ExecutedScraping struct{
	ID int64 `db:"id"`
	ExecutedSpDateTime time.Time `db:"executed_scraping_datetime" json:"executed_scraping_datetime"`
}