package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDbConnection(username string, password string, hostName string, port int) *mongo.Client {
	// mongodb://root:1234@localhost:27017/

	dsn := fmt.Sprintf("mongodb://%v:%v@%v:%v", username, password, hostName, password)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	return client
}
