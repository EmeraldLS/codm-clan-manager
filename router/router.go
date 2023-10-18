package router

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vought-esport-attendance/controller"
)

func Run() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        false,
		AllowOrigins:           []string{"https://vought-esport.vercel.app"},
		AllowMethods:           []string{"POST", "PUT", "GET", "DELETE"},
		AllowHeaders:           []string{"Content-Type"},
		AllowCredentials:       false,
		ExposeHeaders:          []string{},
		MaxAge:                 0,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))
	port := os.Getenv("API_PORT")
	r.POST("/attendance", controller.InitializeDbContent)
	r.GET("/tournaments", controller.GetAllTournament)
	r.GET("/tournaments/:id", controller.GetTournament)
	r.POST("/register", controller.RegisterPlayer)
	r.GET("/lobbies/:id/:day_number", controller.GetLobbyByDay)
	r.GET("/lobby/:id", controller.GetLobbyByID)
	r.GET("/lobby_index/:id", controller.GetLobbyByIndex)
	r.GET("/lobby_players/:id/:day_number/:lobby_id", controller.GetPlayersInALobbby)
	r.GET("/users", controller.GetAllUsers)
	r.GET("/users/:player_id", controller.GetSingleUser)
	r.GET("/player_from_lobby/:id", controller.GetPlayerDetailsFromALobby)
	r.POST("/lobby/:id", controller.CreateLobby)
	r.PUT("/lobby/player/:id", controller.AddPlayerKillsInALobby)

	r.Run("0.0.0.0:" + port)
}
