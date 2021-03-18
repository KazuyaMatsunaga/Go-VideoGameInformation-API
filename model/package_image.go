package model

type PackageImg struct {
	ID            int64  `db:"id"`
	PackageImgURL string `db:"package_image_url" json:"package_image_url"`
}
