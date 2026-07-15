package models

import (
	"time"

	"gorm.io/gorm"
)

type RewardHistory struct {
	gorm.Model

	UserID     string    `json:"userId"`
	Reward     string    `json:"reward"`
	EarnedDate time.Time `json:"create_date"`
}
