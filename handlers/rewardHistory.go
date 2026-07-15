package handlers

import (
	"net/http"

	"nextzy-game-be/config"
	"nextzy-game-be/models"

	"github.com/gin-gonic/gin"
)

func CreateRewardHistory(c *gin.Context) {

	var request models.CreateRewardHistoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	rewardRequirement := map[string]int64{
		"A": 5000,
		"B": 7500,
		"C": 10000,
	}

	requiredPoint, exists := rewardRequirement[request.Reward]

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid reward",
		})
		return
	}

	var user models.User

	if err := config.DB.
		Where("user_id = ?", request.UserId).
		First(&user).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	totalPoint, err := GetUserTotalPoint(request.UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to calculate user point",
		})
		return
	}

	if totalPoint < requiredPoint {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "not enough points",
		})
		return
	}

	var existingReward models.RewardHistory

	err = config.DB.
		Where("user_id = ? AND reward = ?",
			request.UserId,
			request.Reward).
		First(&existingReward).Error

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "reward already claimed",
		})
		return
	}

	rewardHistory := models.RewardHistory{
		UserID: request.UserId,
		Reward: request.Reward,
	}

	if err := config.DB.
		Create(&rewardHistory).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save reward",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "reward claimed successfully",
		"reward":  request.Reward,
	})
}
