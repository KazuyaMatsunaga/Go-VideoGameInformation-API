package model

type PfAndPackImgOfGame struct {
	GameID    int64 `db:"game_id"`
	PfID      int64 `db:"platform_id"`
	PackImgID int64 `db:"package_image_id"`
}
