package model

type PackageImage struct{
	ID int64 `db:"id"`
	PackageImgURL string `db:"package_image_url" json:"package_image_url"`
	ExecutedGetImgRequestID int64 `db:"executed_get_image_request_id" json:"executed_get_image_request_id"`
}