package model

type DatetimeOfFoundPackImg struct {
	PackImgID                 int64 `db:"package_image_id" json:"package_image_id"`
	ExecutedSearchForGetImgID int64 `db:"executed_scraping_for_get_image_id" json:"executed_scraping_for_get_image_id"`
}
