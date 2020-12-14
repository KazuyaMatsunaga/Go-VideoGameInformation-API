package model

import (
	"time"
)

type ExecutedScrpForGetImage struct{
	ID int64 `db:"id"`
	ExecutedScrpForGetImgAt time.Time `db:"executed_scraping_for_get_image_at" json:"executed_scraping_for_get_image_at"`
}