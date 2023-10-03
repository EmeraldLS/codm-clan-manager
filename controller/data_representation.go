package controller

import (
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/vought-esport-attendance/model"
)

func NewLobby(lobbies model.Lobbies, lobbyNumber int) []model.Lobby {
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

func createDay() model.Lobbies {
	return model.Lobbies{
		Lobbies: []model.Lobby{},
	}
}

func RepresentDBData() model.Attendance {
	return model.Attendance{
		Day1: model.Lobbies{
			Lobbies: []model.Lobby{{
				LobbyID:     uuid.NewString(),
				Date:        carbon.Now().ToDateString(),
				LobbyNumber: 1,
				Players: []model.Player{{
					Name:     "Lawrence",
					PlayerID: "VOUGHT_PLAYER_1",
					Kills:    5,
				}},
			}},
		},
		Day2: model.Lobbies{
			Lobbies: []model.Lobby{{
				LobbyID:     uuid.NewString(),
				Date:        carbon.Now().ToDateString(),
				LobbyNumber: 1,
				Players: []model.Player{{
					Name:     "Sanni",
					PlayerID: "VOUGHT_PLAYER_2",
					Kills:    2,
				}},
			}},
		},
		Day3: createDay(),
		Day4: createDay(),
		Day5: createDay(),
	}
}
