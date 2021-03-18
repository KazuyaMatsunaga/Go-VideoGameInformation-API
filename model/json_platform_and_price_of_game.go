package model

type JsonPfAndPriceOfGame struct {
	PfAbbrName string `db:"platform_abbreviation_name" json:"platform"`
	Price      string `db:"price" json:"price"`
}
