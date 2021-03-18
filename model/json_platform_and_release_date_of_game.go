package model

import (
	"time"
)

type JsonPfAndRlDateOfGame struct {
	PfAbbrName string    `db:"platform_abbreviation_name" json:"platform"`
	RlDate     time.Time `db:"release_date" json:"release_date"`
}
