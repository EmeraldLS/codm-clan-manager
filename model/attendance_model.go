package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Attendance struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Day1 Lobbies
	Day2 Lobbies
	Day3 Lobbies
	Day4 Lobbies
	Day5 Lobbies
}

type Lobbies struct {
	Lobbies []Lobby
}

type Lobby struct {
	LobbyID     string `bson:"lobby_id,omitempty"`
	Date        string `bson:"date,omitempty"`
	LobbyNumber int    `json:"lobby_number" bson:"lobby_number" validate:"required"`
	Players     []Player
}

type Player struct {
	PlayerCode int    `json:"layer_code,omitempty" bson:"player_code,omitempty"`
	PlayerID   string `json:"player_id,omitempty" bson:"player_id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Kills      int    `json:"kills,omitempty" bson:"kills,omitempty" validate:"required"`
	UpdatedAt  string `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Day struct {
	DayNumber int `json:"day_number" validate:"required"`
}

type LobbyDetails struct {
	DayNumber int    `json:"day_number"  validate:"required"`
	LobbyID   string `json:"lobby_id"  validate:"required"`
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
