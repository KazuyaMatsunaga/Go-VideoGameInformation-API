package model

type Platform struct{
	ID int64 `db:"id"`
	PfName string `db:"platform_name" json:"platform_name"`
	PfURL string `db:"platform_url" json:"platform_url"`
}