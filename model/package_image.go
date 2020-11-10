package model

type PackageImage struct{
	ID int64 `db:"id"`
	PackageImageURL string `db:"package_image_url" json:"package_image_url"`
	ExecutedGetImageRequestID int64 `db:"executed_get_image_request_id" json:"executed_get_image_request_id"`
}