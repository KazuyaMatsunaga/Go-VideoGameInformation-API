package model

type DatetimeOfFoundGame struct{
	GameID int64 `db:"game_id" json:"game_id"`
	ExecutedScrpForGameInfoID int64 `db:"executed_scraping_for_get_game_information_id" json:"executed_scraping_for_get_game_information_id"`
}