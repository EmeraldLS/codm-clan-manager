package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Attendance struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TournamentName string             `json:"tournament_name,omitempty" bson:"tournament_name,omitempty"`
	Day1           LobbiesData
	Day2           LobbiesData
	Day3           LobbiesData
	Day4           LobbiesData
	Day5           LobbiesData
	Day6           LobbiesData
	Day7           LobbiesData
}

type LobbiesData struct {
	Lobbies []Lobby
}

type Lobby struct {
	LobbyID     string `bson:"lobby_id,omitempty"`
	Date        string `bson:"date,omitempty"`
	LobbyNumber int    `json:"lobby_number" bson:"lobby_number" validate:"required"`
	Players     []Player
}

type Player struct {
	PlayerCode int    `json:"player_code,omitempty" bson:"player_code,omitempty"`
	PlayerID   string `json:"player_id,omitempty" bson:"player_id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Kills      int    `json:"kills,omitempty" bson:"kills,omitempty" validate:"required"`
	UpdatedAt  string `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Day struct {
	DayNumber int `json:"day_number" validate:"required"`
}

type PlayerDetails struct {
	DayNumber int    `json:"day_number"  validate:"required"`
	LobbyID   string `json:"lobby_id"  validate:"required"`
	PlayerID  string `json:"player_id" validate:"required"`
}

type LobbyCreation struct {
	LobbyID     string `bson:"lobby_id,omitempty"`
	Date        string `bson:"date,omitempty"`
	LobbyNumber int    `json:"lobby_number" bson:"lobby_number" validate:"required"`
	DayNumber   int    `json:"day_number,omitempty" bson:"day_number,omitempty" validate:"required"`
	Players     []Player
}

type KillCount struct {
	LobbyID   string `json:"lobby_id,omitempty" validate:"required"`
	DayNumber int    `json:"day_number,omitempty" bson:"day_number,omitempty" validate:"required"`
	PLayerID  string `json:"player_id,omitempty" validate:"required"`
	Kills     int    `json:"kills,omitempty" validate:"required"`
}
