package controller

import (
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/vought-esport-attendance/model"
)

func NewLobby(lobbies model.LobbiesData, lobbyNumber int) []model.Lobby {
	lobby := model.Lobby{
		LobbyNumber: lobbyNumber,
		LobbyID:     uuid.NewString(),
		Date:        carbon.Now().ToDateString(),
	}
	lobbies.Lobbies = append(lobbies.Lobbies, lobby)
	return lobbies.Lobbies
}

func NewPlayer(playerId string, kills int, players []model.Player) []model.Player {
	player := model.Player{
		PlayerID: playerId,
		Kills:    kills,
	}

	players = append(players, player)
	return players
}

func createDay() model.LobbiesData {
	return model.LobbiesData{
		Lobbies: []model.Lobby{},
	}
}

func RepresentDBData(tournament_name string) model.Attendance {
	attendance := model.Attendance{
		TournamentName: tournament_name,
		Day1:           createDay(),
		Day2:           createDay(),
		Day3:           createDay(),
		Day4:           createDay(),
		Day5:           createDay(),
		Day6:           createDay(),
		Day7:           createDay(),
	}

	return attendance
}
