package model

type DatetimeOfFoundGame struct{
	ID int64 `db:"id"`
	ExecutedScrpForGameInfoID int64 `db:"executed_scraping_for_get_game_information_id" json:"executed_scraping_for_get_game_information_id"`
}