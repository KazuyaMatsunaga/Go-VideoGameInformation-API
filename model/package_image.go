package model

import (
	"time"
)


type PackageImage struct{
	ID int64 `db:"id"`
	PackageImgURL string `db:"package_image_url" json:"package_image_url"`
	CreatedTime      time.Time   `db:"ctime" json:"created_time"`
	UpdatedTime      time.Time   `db:"utime" json:"updated_time"`
}