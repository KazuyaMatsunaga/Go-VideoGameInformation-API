package model

type PfAndPriceOfGame struct{
	GameID int64 `db:"game_id"`
	PfID int64 `db:"platform_id"`
	Price string `db:"price"`
}