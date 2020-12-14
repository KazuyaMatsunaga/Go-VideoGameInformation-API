package model

import (
	"time"
)

type Platform struct{
	ID int64 `db:"id"`
	PfAbbrName string `db:"platform_abbreviation_name" json:"platform_abbreviation_name"`
	PfName string `db:"platform_name" json:"platform_name"`
	CreatedTime      time.Time   `db:"ctime" json:"created_time"`
	UpdatedTime      time.Time   `db:"utime" json:"updated_time"`
}