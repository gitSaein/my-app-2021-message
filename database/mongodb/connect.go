package mongodb

import (
	"context"
	"time"

	config "my-app-2021-message/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnInfo struct {
	Conn       *mongo.Client
	Ctx        context.Context
	Config     config.Config
	Cancel     context.CancelFunc
	Collection *mongo.Collection
	Err        error
}

func init() {

}

func Conn(env string) *ConnInfo {
	config := config.GetCongif(env)

	credentials := options.Credential{
		Username: config.Database.MongoDB.User,
		Password: config.Database.MongoDB.Pwd,
	}
	clientOptions := options.
		Client().
		ApplyURI(config.Database.MongoDB.Uri).
		SetAuth(credentials)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return &ConnInfo{client, ctx, config, cancel, nil, err}
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return &ConnInfo{client, ctx, config, cancel, nil, err}
	}

	collection := client.
		Database(config.Database.MongoDB.MessageDatabase).
		Collection(config.Database.MongoDB.MessageCollection)
	return &ConnInfo{client, ctx, config, cancel, collection, nil}

}
