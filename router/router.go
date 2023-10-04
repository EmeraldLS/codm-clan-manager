package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vought-esport-attendance/controller"
)

func Run() {
	r := gin.Default()
	port := os.Getenv("API_PORT")
	r.POST("/attendance", controller.InitializeDbContent)
	r.POST("/register", controller.RegisterPlayer)
	r.GET("/lobbies/:id", controller.GetLobbyByDay)
	r.GET("/lobby/:id", controller.GetLobbyByID)
	r.GET("/lobby_index/:id", controller.GetLobbyByIndex)
	r.GET("/lobby_players/:id", controller.GetPlayersInALobbby)
	r.GET("/player_from_lobby/:id", controller.GetPlayerDetailsFromALobby)
	r.POST("/lobby/:id", controller.CreateLobby)
	r.POST("/lobby/player/:id", controller.AddPlayerKillsInALobby)

	r.Run("0.0.0.0:" + port)
}
