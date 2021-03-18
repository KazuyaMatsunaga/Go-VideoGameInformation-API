package model

type JsonPfAndPackImgOfGame struct {
	PfAbbrName    string `db:"platform_abbreviation_name" json:"platform"`
	PackageImgURL string `db:"package_image_url" json:"package_image_url"`
}
