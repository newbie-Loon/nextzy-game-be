package models

type CreatePointHistoryRequest struct {
	UserId string `json:"userId"`
	Point  int    `json:"point"`
}
