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

func mongoConn(env string) (*mongo.Client, context.Context, context.CancelFunc, config.Config, error) {
	var err error
	config := config.GetCongif(env)

	credentials := options.Credential{
		Username: config.Database.MongoDB.User,
		Password: config.Database.MongoDB.Pwd,
	}
	clientOptions := options.Client().
		ApplyURI(config.Database.MongoDB.Uri).
		SetAuth(credentials)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		err = err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		err = err
	}

	log.Println("MongoDB Connected")

	return client, ctx, cancel, config, err

}
