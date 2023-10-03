package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/vought-esport-attendance/config"
	"github.com/vought-esport-attendance/model"
)

func InitializeDbContent(c *gin.Context) {
	attendance := RepresentDBData()

	if err := config.InitializeDbContent(attendance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "content created.",
	})
}

func GetLobbyByDay(c *gin.Context) {
	id := c.Param("id")
	var day model.Day

	if err := c.ShouldBindJSON(&day); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}
	validate := validator.New()
	if err := validate.Struct(day); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	dayLobbbies, err := config.GetAllLobbyInADay(id, day.DayNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, dayLobbbies)
}

func GetLobbyByID(c *gin.Context) {
	id := c.Param("id")
	var lobbyDetails model.LobbyDetails
	if err := c.ShouldBindJSON(&lobbyDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	if err := validate.Struct(lobbyDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	lobby, err := config.GetLobbyByID(id, lobbyDetails.LobbyID, lobbyDetails.DayNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, lobby)
}

func GetPlayersInALobbby(c *gin.Context) {
	id := c.Param("id")
	var lobbyDetails model.LobbyDetails
	if err := c.ShouldBindJSON(&lobbyDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	if err := validate.Struct(lobbyDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	players, err := config.GetPlayersInALobbby(id, lobbyDetails.LobbyID, lobbyDetails.DayNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, players)

}

func GetPlayerDetailsFromALobby(c *gin.Context) {
	id := c.Param("id")
	var playerDetails model.PlayerDetails
	if err := c.ShouldBindJSON(&playerDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	if err := validate.Struct(playerDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	player, err := config.GetAPlayerFromALobby(id, playerDetails.LobbyID, playerDetails.PlayerID, playerDetails.DayNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, player)
}

// func InsertPlayerKills(c *gin.Context) {
// 	playerID := c.Param("player_id")
// 	user, err := config.GetUserByID(playerID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"response": err.Error(),
// 		})
// 		c.Abort()
// 		return
// 	}

// 	player := model.Player{
// 		PlayerID: user.PlayerID,
// 		Name:     user.Name,
// 		Kills:    6,
// 	}
// }

func CreateLobby(c *gin.Context) {
	_id := c.Param("id")
	var lobby model.LobbyCreation
	if err := c.ShouldBindJSON(&lobby); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	validate := validator.New()
	if err := validate.Struct(lobby); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}

	lobby.Date = carbon.Now().ToDateString()
	lobby.LobbyID = uuid.NewString()

	// lobby := model.Lobby{
	// 	LobbyID:     uuid.NewString(),
	// 	LobbyNumber: 1,
	// 	Date:        carbon.Now().ToDateString(),
	// 	Players:     []model.Player{},
	// }
	allLobby, err := config.CreateLobby(_id, lobby)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, allLobby)
}

// func UpdateKillCount(c *gin.Context) {

// 	lobbyID := c.Param("lobby_id")
// 	playerID := c.Param("player_id")

// 	var killCount model.KillCount

// 	if err := c.ShouldBindJSON(&killCount); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"response": err.Error(),
// 		})
// 		c.Abort()
// 		return
// 	}

// 	var validate = validator.New()
// 	if err := validate.Struct(killCount); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"response": err.Error(),
// 		})
// 		c.Abort()
// 		return
// 	}

// 	lobby, err := config.CheckLobbyID(lobbyID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"response": err.Error(),
// 		})
// 		c.Abort()
// 		return
// 	}

// 	player, err := config.CheckPlayerID(playerID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"response": err.Error(),
// 		})
// 		c.Abort()
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"Lobby":  lobby,
// 		"Player": player,
// 	})

// }
