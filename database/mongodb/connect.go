package mongodb

import (
	"context"
	"errors"
	"fmt"
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

const (
	CHAT    = "chat"
	MESSAGE = "message"
)

func connect(env string) *ConnInfo {
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
		return &ConnInfo{
			Conn: client, Ctx: ctx, Config: config, Cancel: cancel, Err: err}
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return &ConnInfo{
			Conn: client, Ctx: ctx, Config: config, Cancel: cancel, Err: err}
	}
	return &ConnInfo{
		Conn: client, Ctx: ctx, Config: config, Cancel: cancel}
}

func ApproachCollection(env string, collection string) *ConnInfo {
	client := connect(env)

	switch collection {
	case CHAT:
		client.Collection = client.Conn.
			Database(client.Config.Database.MongoDB.Database).
			Collection(client.Config.Database.MongoDB.ChatCollection)

		return client
	case MESSAGE:
		client.Collection = client.Conn.
			Database(client.Config.Database.MongoDB.Database).
			Collection(client.Config.Database.MongoDB.MessageCollection)

		return client
	default:
		client.Err = errors.New(fmt.Sprintf("Not Found Collection: %s", collection))

		return client
	}

}
