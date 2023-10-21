package config

import (
	"context"
	"time"

	"github.com/vought-esport-attendance/db"
	"github.com/vought-esport-attendance/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var PlayersCollection = db.PlayersCollection

func CheckPlayersID(ids []string) ([]model.User, string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	var players []model.User
	for _, id := range ids {
		var user model.User
		filter := bson.M{"player_id": id}
		err := PlayersCollection.FindOne(ctx, filter).Decode(&user)
		defer cancel()
		if err != nil {
			return []model.User{}, id, err
		}
		players = append(players, user)
	}

	return players, "", nil
}

func GetUserByID(playerID string) (model.User, error) {
	filter := bson.M{"player_id": playerID}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	var user model.User
	if err := PlayersCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return model.User{}, err
	}

	return user, nil

}

func GetAllUsers(page int64) ([]model.User, error) {
	filter := bson.M{}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	findOptions := options.Find()

	findOptions.SetSort(bson.M{"player_code": -1})
	if page != 0 {
		var perpage int64 = 1000
		findOptions.SetSkip((page - 1) * perpage)
		findOptions.SetLimit(perpage)
	}

	cursor, err := PlayersCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return []model.User{}, err
	}
	var users []model.User
	for cursor.Next(ctx) {
		var user model.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

func GetSingleUser(playerId string) (model.User, error) {
	filter := bson.M{"player_id": playerId}
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	var user model.User
	if err := PlayersCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return model.User{}, err
	}
	return user, nil
}
