package config

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/vought-esport-attendance/db"
	"github.com/vought-esport-attendance/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var AttendanceCollection = db.AttendanceCollection

func InitializeDbContent(attedance model.Attendance) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	_, err := AttendanceCollection.InsertOne(ctx, attedance)
	defer cancel()
	if err != nil {
		return err
	}
	defer cancel()
	return nil
}

func RegisterPlayer(player model.User) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	_, err := PlayersCollection.InsertOne(ctx, player)
	defer cancel()
	if err != nil {
		return err
	}
	defer cancel()
	return nil
}

func GetAllTournament() ([]model.Attendance, error) {
	filter := bson.M{}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	// findOptions := options.Find()

	cursor, err := AttendanceCollection.Find(ctx, filter)
	if err != nil {
		return []model.Attendance{}, err
	}
	var tournaments []model.Attendance
	for cursor.Next(ctx) {
		var tournament model.Attendance
		cursor.Decode(&tournament)
		tournaments = append(tournaments, tournament)
	}

	return tournaments, nil
}

func GetTournament(id string) (model.Attendance, error) {
	_id, err := ConvertStringToOBjectID(id)
	if err != nil {
		return model.Attendance{}, err
	}

	filter := bson.M{"_id": _id}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	var tournament model.Attendance
	err = AttendanceCollection.FindOne(ctx, filter).Decode(&tournament)
	if err != nil {
		return model.Attendance{}, err
	}

	return tournament, nil
}

func GetAllLobbyInADay(id string, day int) ([]model.Lobby, error) {
	attedance, err := getAttedanceByID(id)
	if err != nil {
		return []model.Lobby{}, err
	}

	var allLobby []model.Lobby

	switch day {
	case 1:
		allLobby = attedance.Day1.Lobbies
	case 2:
		allLobby = attedance.Day2.Lobbies
	case 3:
		allLobby = attedance.Day3.Lobbies
	case 4:
		allLobby = attedance.Day4.Lobbies
	case 5:
		allLobby = attedance.Day5.Lobbies
	case 6:
		allLobby = attedance.Day6.Lobbies
	case 7:
		allLobby = attedance.Day7.Lobbies
	default:
		return []model.Lobby{}, errors.New("invalid day number")
	}

	return allLobby, nil
}

func GetTotalPlayerKillsInADay(id, playerID string, day int) (int, error) {
	allLobby, err := GetAllLobbyInADay(id, day)
	total := 0
	if err != nil {
		return total, err
	}
	for _, lobby := range allLobby {
		for _, player := range lobby.Players {
			if player.PlayerID == playerID {
				total += player.Kills
			}
		}
	}

	return total, nil
}

func GetTotalPlayerKillsInWholeTournament(id, playerID string) (int, error) {
	total := 0

	for day := 1; day <= 7; day++ {
		dayTotal, err := GetTotalPlayerKillsInADay(id, playerID, day)
		if err != nil {
			return total, err
		}
		fmt.Printf("Day %v, you've %v kills\n", day, dayTotal)
		total += dayTotal
	}
	return total, nil
}

func getAttedanceByID(id string) (model.Attendance, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Attendance{}, err
	}
	filter := bson.M{"_id": _id}
	var attedance model.Attendance
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := AttendanceCollection.FindOne(ctx, filter).Decode(&attedance); err != nil {
		return model.Attendance{}, err
	}
	return attedance, nil
}

func GetLobbyByID(id string, lobbyID string, day int) (model.Lobby, error) {
	allLobbies, err := GetAllLobbyInADay(id, day)
	if err != nil {
		return model.Lobby{}, err
	}
	var foundLobby model.Lobby
	for _, lobby := range allLobbies {
		if lobby.LobbyID == lobbyID {
			foundLobby = lobby
		}
	}
	return foundLobby, nil
}

func GetPlayersInALobbby(id string, lobbyID string, day int) ([]model.Player, error) {
	lobby, err := GetLobbyByID(id, lobbyID, day)
	if err != nil {
		return []model.Player{}, err
	}

	return lobby.Players, nil

}

func GetAPlayerFromALobby(id string, lobbyID string, playerID string, day int) (model.Player, error) {
	players, err := GetPlayersInALobbby(id, lobbyID, day)
	if err != nil {
		return model.Player{}, err
	}
	var foundPlayer model.Player
	for _, player := range players {
		if player.PlayerID == playerID {
			foundPlayer = player
		}
	}
	return foundPlayer, nil
}

func ConvertStringToOBjectID(id string) (primitive.ObjectID, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return _id, nil
}

func CreateLobby(id string, lobbyCreation model.LobbyCreation) ([]model.Lobby, error) {

	lobby := model.Lobby{
		LobbyID:     lobbyCreation.LobbyID,
		LobbyNumber: lobbyCreation.LobbyNumber,
		Date:        lobbyCreation.Date,
		Players:     []model.Player{},
	}

	if lobby.LobbyNumber > 3 || lobby.LobbyNumber < 1 {
		return []model.Lobby{}, errors.New("invalid lobby number")
	}

	allLobby, err := GetAllLobbyInADay(id, lobbyCreation.DayNumber)
	if err != nil {
		return []model.Lobby{}, err
	}

	for _, aLobby := range allLobby {
		if aLobby.LobbyNumber == lobby.LobbyNumber {
			return []model.Lobby{}, errors.New("a lobby with provided number already exist")
		} else if len(allLobby) == 3 {
			return []model.Lobby{}, errors.New("you can't create anymore lobby today")
		}
	}

	allLobby = append(allLobby, lobby)
	var updateObj = bson.M{}
	if lobbyCreation.DayNumber > 7 || lobbyCreation.DayNumber < 1 {
		return []model.Lobby{}, errors.New("invalid day number")
	}
	queryString := fmt.Sprintf("day%v.lobbies", lobbyCreation.DayNumber)
	updateObj[queryString] = allLobby
	update := bson.M{"$set": updateObj}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	_id, err := ConvertStringToOBjectID(id)
	if err != nil {
		return []model.Lobby{}, err
	}
	filter := bson.M{"_id": _id}
	_, err = AttendanceCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return []model.Lobby{}, err
	}

	return allLobby, nil
}

func InsertPlayerKillInALobby(id string, lobbyID string, playerCreation model.KillCount, day int) ([]model.Player, error) {
	allPlayers, err := GetPlayersInALobbby(id, lobbyID, day)
	if err != nil {
		return []model.Player{}, err
	}

	user, err := GetUserByID(playerCreation.PLayerID)
	if err != nil {
		return []model.Player{}, err
	}

	player := model.Player{
		PlayerCode: user.PlayerCode,
		PlayerID:   user.PlayerID,
		Name:       user.Name,
		Kills:      playerCreation.Kills,
	}

	var updateObj = bson.M{}
	queryString := fmt.Sprintf("day%v.lobbies", day)

	allLobby := []model.Lobby{}
	var lobbyErr error = nil
	switch day {
	case 1:
		allLobby, lobbyErr = GetAllLobbyInADay(id, 1)
	case 2:
		allLobby, lobbyErr = GetAllLobbyInADay(id, 2)
	case 3:
		allLobby, lobbyErr = GetAllLobbyInADay(id, 3)
	case 4:
		allLobby, lobbyErr = GetAllLobbyInADay(id, 4)
	case 5:
		allLobby, lobbyErr = GetAllLobbyInADay(id, 5)
	case 6:
		allLobby, lobbyErr = GetAllLobbyInADay(id, 6)
	case 7:
		allLobby, lobbyErr = GetAllLobbyInADay(id, 7)
	}

	if lobbyErr != nil {
		return []model.Player{}, lobbyErr
	}

	var newLobbies = []model.Lobby{}

	for _, aLobby := range allLobby {
		if aLobby.LobbyID == lobbyID {
			for _, aPlayer := range aLobby.Players {
				if aPlayer.PlayerID == player.PlayerID {
					return []model.Player{}, errors.New("player already exist in this lobby")
				}
			}
			allPlayers = append(allPlayers, player)
			aLobby.Players = allPlayers

		}
		if len(aLobby.Players) > 4 {
			return []model.Player{}, errors.New("lobby is filled up")
		}

		newLobbies = append(newLobbies, aLobby)
	}

	updateObj[queryString] = newLobbies
	update := bson.M{"$set": updateObj}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	_id, err := ConvertStringToOBjectID(id)
	if err != nil {
		return []model.Player{}, err
	}
	filter := bson.M{"_id": _id}
	_, err = AttendanceCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return []model.Player{}, err
	}
	players, err := GetPlayersInALobbby(id, lobbyID, day)
	if err != nil {
		return []model.Player{}, err
	}

	return players, nil

}

func GetLobbyByIndex(id string) error {
	_id, err := ConvertStringToOBjectID(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": _id}
	projection := bson.M{"lobby": bson.M{
		"$arrayElemAt": []interface{}{"$day1.lobbies", 1},
	}}

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := AttendanceCollection.Aggregate(ctx, mongo.Pipeline{
		{{"$match", filter}},
		{{"$project", projection}},
	})
	if err != nil {
		return err
	}

	for cursor.Next(ctx) {
		var result = bson.M{}
		cursor.Decode(&result)
		key := result["_id"]
		fmt.Println(key)
	}
	return nil
}

/*

	Below was the first secnario I used. Keeping it for future references.

*/

// func InsertLobby(id string) (model.Attendance, error) {

// }

// func CheckLobbyID(lobbyID string) (model.Lobby, error) {
// 	var filters = []bson.M{}
// 	numDays := 5
// 	numLobbies := 3
// 	for dayNum := 1; dayNum <= numDays; dayNum++ {
// 		for lobbyNum := 1; lobbyNum <= numLobbies; lobbyNum++ {
// 			filterCondition := bson.M{
// 				fmt.Sprintf("day%v.lobby%v.lobby_id", dayNum, lobbyNum): lobbyID,
// 			}
// 			filters = append(filters, filterCondition)
// 		}
// 	}

// 	pipeline := mongo.Pipeline{
// 		{
// 			{
// 				"$match", bson.M{
// 					"$or": filters,
// 				},
// 			},
// 		},
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	cursor, err := AttendanceCollection.Aggregate(ctx, pipeline)
// 	defer cancel()
// 	if err != nil {
// 		return model.Lobby{}, err
// 	}

// 	var results []model.Lobby
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		var res model.Lobby
// 		cursor.Decode(&res)
// 		results = append(results, res)
// 	}

// 	if len(results) == 0 {
// 		return model.Lobby{}, errors.New("no Lobby with provided ID found")
// 	}

// 	var lobby model.Lobby

// 	fmt.Println(results)

// 	for _, result := range results {
// 		lobby = result
// 	}
// 	return lobby, nil
// }

// func CheckPlayerID(playerID string) (model.Player, error) { //It checks player Id from the attendance collection
// 	filters := []bson.M{}

// 	var results []model.Player
// 	var player model.Player

// 	numDays := 5
// 	numLobbies := 3
// 	numPlayers := 5

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	for day := 1; day <= numDays; day++ {
// 		for lobby := 1; lobby <= numLobbies; lobby++ {
// 			for player := 1; player < numPlayers; player++ {
// 				filterCondition := bson.M{fmt.Sprintf("day%v.lobby%v.player%v.player_id", day, lobby, player): playerID}
// 				filters = append(filters, filterCondition)
// 			}
// 		}
// 	}

// 	// fmt.Printf("\n\nFilter: %v\n\n", filter)

// 	pipeline := mongo.Pipeline{
// 		{
// 			{
// 				"$match", bson.M{
// 					"$or": filters,
// 				},
// 			},
// 		},
// 	}

// 	cursor, err := AttendanceCollection.Aggregate(ctx, pipeline)
// 	defer cancel()
// 	defer cursor.Close(ctx)
// 	if err != nil {
// 		return model.Player{}, err
// 	}
// 	for cursor.Next(ctx) {
// 		var result model.Player
// 		cursor.Decode(&result)
// 		results = append(results, result)
// 	}

// 	for _, result := range results {
// 		player = result
// 	}
// 	return player, nil
// }
