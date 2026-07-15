package models

type CreateRewardHistoryRequest struct {
	UserId string `json:"userId"`
	Reward string `json:"reward"`
}
