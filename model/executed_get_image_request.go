package model

import (
	"time"
)

type ExecutedGetImageRequest struct{
	ID int64 `db:"id"`
	ExecutedGetImageRequestDateTime time.Time `db:"executed_get_image_request_datetime" json:"executed_get_image_request_datetime"`
}