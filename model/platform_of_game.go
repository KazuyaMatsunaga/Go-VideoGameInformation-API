package model

type PfOfGame struct{
	GameID int64 `db:"game_id" json:"game_id"`
	PfID int64 `db:"platform_id" json:"platform_id"`
}