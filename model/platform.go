package model

type Platform struct{
	ID int64 `db:"id" json:"id"`
	PfAbbrName string `db:"platform_abbreviation_name" json:"platform_abbreviation_name"`
	PfName string `db:"platform_name" json:"platform_name"`
}