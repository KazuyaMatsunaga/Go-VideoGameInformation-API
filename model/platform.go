package model

type Platform struct{
	ID int64 `db:"id"`
	PlatformName string `db:"platform_name" json:"platform_name"`
}