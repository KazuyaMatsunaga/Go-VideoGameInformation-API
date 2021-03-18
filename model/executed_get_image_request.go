package model

import (
	"time"
)

type ExecutedSearchForGetImage struct {
	ID                        int64     `db:"id"`
	ExecutedSearchForGetImgAt time.Time `db:"executed_scraping_for_get_image_at" json:"executed_scraping_for_get_image_at"`
}
