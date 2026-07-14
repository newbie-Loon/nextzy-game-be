package models

import (
	"time"

	"gorm.io/gorm"
)

type PointHistory struct {
	gorm.Model

	UserID     string    `json:"guestId"`
	Point      int       `json:"point"`
	EarnedDate time.Time `json:"create_date"`
}
