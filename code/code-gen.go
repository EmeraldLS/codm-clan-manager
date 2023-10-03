package code

import (
	"context"
	"fmt"
	"time"

	"github.com/vought-esport-attendance/db"
	"github.com/vought-esport-attendance/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMaxPlayerCode() int {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	filter := bson.M{}
	findOptions := options.Find().SetSort(bson.M{"player_code": -1}).SetLimit(1)
	cursor, _ := db.PlayersCollection.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)
	var players []model.User
	for cursor.Next(ctx) {
		var player model.User
		cursor.Decode(&player)
		players = append(players, player)
	}
	var maxCode int
	for _, player := range players {
		maxCode = player.PlayerCode
	}
	return maxCode
}

func GenPlayerID(player_code int) string {
	prefix := "VOUGHT_PLAYER_"
	userID := fmt.Sprintf("%v%d", prefix, player_code)
	return userID
}
