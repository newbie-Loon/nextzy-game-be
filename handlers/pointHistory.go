package handlers

import (
	"net/http"
	"time"

	"nextzy-game-be/config"
	"nextzy-game-be/models"

	"github.com/gin-gonic/gin"
)

func CreatePointHistory(c *gin.Context) {
	var request models.CreatePointHistoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	allowedPoints := map[int]bool{
		300:  true,
		500:  true,
		1000: true,
		3000: true,
	}

	if !allowedPoints[request.Point] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid point value",
		})
		return
	}

	tx := config.DB.Begin()

	var user models.User

	if err := tx.
		Where("user_id = ?", request.UserId).
		First(&user).Error; err != nil {

		tx.Rollback()

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	history := models.PointHistory{
		UserID:     request.UserId,
		Point:      request.Point,
		EarnedDate: time.Now(),
	}

	if err := tx.Create(&history).Error; err != nil {
		tx.Rollback()

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create point history",
		})
		return
	}
	tx.Commit()

	var totalPoint int64

	totalPoint, err := GetUserTotalPoint(request.UserId)
	finalPoint := totalPoint
	if finalPoint > 10000 {

		finalPoint = 10000

	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get total points",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"earnedPoint": request.Point,
		"totalPoint":  finalPoint,
		"message":     "point added successfully",
	})
}
