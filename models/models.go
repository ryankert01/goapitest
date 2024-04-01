package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type AD struct {
	gorm.Model
	Title      string       `json:"title" gorm:"type:text"`
	StartAt    time.Time    `json:"startAt" gorm:"type:timestamp"`
	EndAt      time.Time    `json:"endAt" gorm:"type:timestamp"`
	IsActive   bool         `gorm:"default:False;type:boolean"`
	Conditions []Conditions `json:"conditions"`
}

type Conditions struct {
	gorm.Model
	ADID     uint
	AgeStart int            `json:"ageStart" gorm:"default:0;type:integer"`
	AgeEnd   int            `json:"ageEnd" gorm:"default:100;type:integer"`
	Gender   pq.StringArray `json:"gender" gorm:"type:text[]"`
	Country  pq.StringArray `json:"country" gorm:"type:text[]"`
	Platform pq.StringArray `json:"platform" gorm:"type:text[]"`
}
