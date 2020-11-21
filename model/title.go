package model

type Title struct{
	ID int64 `db:"id"`
	TitleName string `db:"title_name" json:"title_name"`
	TitleURL string `db:"title_url" json:"title_url"`
}