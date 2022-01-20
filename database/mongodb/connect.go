package mongodb

import (
	"context"
	"log"
	"time"

	config "my-app-2021-message/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {

}

func mongoConn(env string) (*mongo.Client, context.Context, context.CancelFunc, config.Config) {

	config := config.GetCongif(env)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	clientOptions := options.Client().ApplyURI(config.Database.Uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB Connected")

	return client, ctx, cancel, config

}
