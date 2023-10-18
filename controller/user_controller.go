package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-module/carbon"
	"github.com/vought-esport-attendance/code"
	"github.com/vought-esport-attendance/config"
	"github.com/vought-esport-attendance/model"
)

func RegisterPlayer(c *gin.Context) {
	var playerDetails model.User
	if err := c.ShouldBindJSON(&playerDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	if err := validate.Struct(&playerDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}
	maxCode := code.GetMaxPlayerCode()
	playerDetails.PlayerCode = maxCode + 1
	playerDetails.PlayerID = code.GenPlayerID(playerDetails.PlayerCode)
	playerDetails.RegisteredAt = carbon.Now().ToDateTimeString()
	playerDetails.UpdatedAt = carbon.Now().ToDateTimeString()

	if err := config.RegisterPlayer(playerDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "success",
		"User":     playerDetails,
	})

}

func GetAllUsers(c *gin.Context) {
	users, err := config.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, users)

}

func GetSingleUser(c *gin.Context) {
	playerId := c.Param("player_id")
	user, err := config.GetSingleUser(playerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, user)
}
