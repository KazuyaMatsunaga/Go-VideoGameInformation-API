package model

type DatetimeOfFoundPackImg struct{
	ID int64 `db:"id"`
	ExecutedSearchForGetImgID int64 `db:"executed_scraping_for_get_image_id" json:"executed_scraping_for_get_image_id"`
}