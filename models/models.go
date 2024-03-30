package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AD struct {
	gorm.Model
	Title      string         `json:"title"`
	StartAt    time.Time      `json:"startAt"`
	EndAt      time.Time      `json:"endAt"`
	Conditions datatypes.JSON `json:"conditions"`
}
