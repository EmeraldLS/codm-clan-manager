package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vought-esport-attendance/controller"
	"github.com/vought-esport-attendance/helpers"
	"github.com/vought-esport-attendance/middleware"
)

func Run() {
	r := gin.Default()
	clientOriginUrl := helpers.SafeGetEnv("CLIENT_ORIGIN_URL")
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://vought-esports.netlify.app", "http://localhost:3000", "https://voughtesport.lawrencesegun.xyz", clientOriginUrl},
		AllowMethods: []string{"POST", "PUT", "GET", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	port := helpers.SafeGetEnv("API_PORT")

	audience := helpers.SafeGetEnv("AUTH0_AUDIENCE")
	domain := helpers.SafeGetEnv("AUTH0_DOMAIN")

	securedMiddleware := middleware.ValidateJWT(audience, domain)

	secure := r.Group("/")
	secure.Use(securedMiddleware)

	secure.POST("/attendance", controller.InitializeDbContent)
	secure.POST("/register", controller.RegisterPlayer)
	secure.POST("/lobby/:id", controller.CreateLobby)

	r.GET("/tournaments", controller.GetAllTournament)
	r.GET("/tournaments/:id", controller.GetTournament)
	r.GET("/lobbies/:id/:day_number", controller.GetLobbyByDay)
	r.GET("/lobby/:id", controller.GetLobbyByID)
	// r.GET("/lobby_index/:id", controller.GetLobbyByIndex)
	r.GET("/lobby_players/:id/:day_number/:lobby_id", controller.GetPlayersInALobbby)
	r.GET("/users", controller.GetAllUsers)
	r.GET("/users/all", controller.GetAllUsersWithoutPagination)

	r.GET("/users/:player_id", controller.GetSingleUser)
	r.GET("/player_from_lobby/:id", controller.GetPlayerDetailsFromALobby)
	r.GET("/players/:id/:day_number", controller.GetAllPlayersInAday)

	r.GET("/player_kills_total_day/:id/:day_number/:player_id", controller.GetTotalPlayerKillsInADay)
	r.GET("/player_total_kills_tournament/:id/:player_id", controller.GetTotalPlayerKillsInWholeTournament)

	secure.PUT("/lobby/player/:id", controller.AddPlayerKillsInALobby)
	r.PUT("/user/:player_id", controller.UpdateUserName)

	r.Run("0.0.0.0:" + port)
}
