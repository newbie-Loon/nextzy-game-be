package handlers

import (
	"fmt"
	"net/http"

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
