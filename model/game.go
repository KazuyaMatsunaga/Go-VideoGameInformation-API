package model

type Game struct {
	ID int64 `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Platforms []string `json:"platforms"`
	Genres []string `json:"genres"`
	RlDates []JsonPfAndRlDateOfGame `json:"release_dates"`
	Prices []JsonPfAndPriceOfGame `json:"prices"`
	PackageImgURLs []JsonPfAndPackImgOfGame `json:"package_image_urls"`
}