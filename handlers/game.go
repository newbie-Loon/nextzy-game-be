package handlers

import (
	"net/http"

	"nextzy-game-be/config"
	"nextzy-game-be/models"

	"github.com/gin-gonic/gin"
)

func ResetGameProgress(c *gin.Context) {
	var request models.ResetProgressRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx := config.DB.Begin()

	// Delete play history
	if err := tx.
		Where("user_id = ?", request.UserId).
		Delete(&models.PointHistory{}).Error; err != nil {
		tx.Rollback()

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete play history",
		})
		return
	}

	// Delete reward history
	if err := tx.
		Where("user_id = ?", request.UserId).
		Delete(&models.RewardHistory{}).Error; err != nil {

		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete reward history",
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message": "progress reset successfully",
	})
}
