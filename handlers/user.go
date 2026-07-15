package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"nextzy-game-be/config"
	"nextzy-game-be/models"
)

func CreateUser(c *gin.Context) {
	var user models.User
	fmt.Printf("%+v\n", user)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var existingUser models.User

	err := config.DB.
		Where("user_id = ?", user.UserId).
		First(&existingUser).Error

	if err == nil {
		fmt.Printf("test")
		c.JSON(http.StatusCreated, user)
		return
	}

	config.DB.Create(&user)

	c.JSON(http.StatusCreated, user)
}

func GetUserData(c *gin.Context) {
	userId := c.Param("userId")

	var user models.User
	var playHistory []models.PointHistory
	var rewardHistory []models.RewardHistory

	if err := config.DB.
		Where("user_id = ?", userId).
		First(&user).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	var totalPoint int64

	config.DB.
		Model(&models.PointHistory{}).
		Where("user_id = ?", userId).
		Select("COALESCE(SUM(point), 0)").
		Scan(&totalPoint)

	config.DB.
		Where("user_id = ?", userId).
		Order("earned_date desc").
		Find(&playHistory)

	config.DB.
		Where("user_id = ?", userId).
		Order("created_at desc").
		Find(&rewardHistory)

	rewards := RewardStatus{}

	for _, reward := range rewardHistory {
		switch reward.Reward {
		case "A":
			rewards.A = true
		case "B":
			rewards.B = true
		case "C":
			rewards.C = true
		}
	}

	response := HomeResponse{
		User: UserInfo{
			UserId: user.UserId,
			Point:  int(totalPoint),
		},
		Rewards:       rewards,
		PlayHistory:   convertPlayHistory(playHistory),
		RewardHistory: convertRewardHistory(rewardHistory),
	}

	c.JSON(http.StatusOK, response)
}

type HomeResponse struct {
	User          UserInfo           `json:"user"`
	Rewards       RewardStatus       `json:"rewards"`
	PlayHistory   []PlayHistoryDto   `json:"playHistory"`
	RewardHistory []RewardHistoryDto `json:"rewardHistory"`
}

type UserInfo struct {
	UserId string `json:"userId"`
	Point  int    `json:"point"`
}

type RewardStatus struct {
	A bool `json:"a"`
	B bool `json:"b"`
	C bool `json:"c"`
}

type PlayHistoryDto struct {
	Point      int       `json:"point"`
	EarnedDate time.Time `json:"earnedDate"`
}

type RewardHistoryDto struct {
	Reward    string    `json:"reward"`
	CreatedAt time.Time `json:"createdAt"`
}

func convertPlayHistory(
	items []models.PointHistory,
) []PlayHistoryDto {

	result := make([]PlayHistoryDto, 0)

	for _, item := range items {
		result = append(result, PlayHistoryDto{
			Point:      item.Point,
			EarnedDate: item.EarnedDate,
		})
	}

	return result
}

func convertRewardHistory(
	items []models.RewardHistory,
) []RewardHistoryDto {

	result := make([]RewardHistoryDto, 0)

	for _, item := range items {
		result = append(result, RewardHistoryDto{
			Reward:    item.Reward,
			CreatedAt: item.CreatedAt,
		})
	}

	return result
}

func GetUserTotalPoint(userId string) (int64, error) {
	var total int64

	err := config.DB.
		Model(&models.PointHistory{}).
		Where("user_id = ?", userId).
		Select("COALESCE(SUM(point), 0)").
		Scan(&total).Error

	return total, err
}
