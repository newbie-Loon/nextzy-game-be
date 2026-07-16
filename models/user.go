package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Guest  bool   `json:"guest"`
	UserId string `json:"userId" gorm:"default:gen_random_uuid()"`
}
