package model

import (
	"time"
)

type Game struct {
	ID int64 `db:"id"`
	Title string `db:"title" json:"title"`
	CreatedTime      time.Time   `db:"ctime" json:"created_time"`
	UpdatedTime      time.Time   `db:"utime" json:"updated_time"`
}