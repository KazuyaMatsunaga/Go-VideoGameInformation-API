package model

import (
	"time"
)

type ExecutedScrpForGameInfo struct {
	ID                           int64     `db:"id"`
	ExecutedScrpingForGameInfoAt time.Time `db:"executed_scraping_for_get_game_information_at" json:"executed_scraping_for_get_game_information_at"`
}
