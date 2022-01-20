package mongodb

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const message_database = "local"
const message_collection = "message"

func InsertMessage(env string) {
	client, ctx, cancel, config := mongoConn(env)
	defer client.Disconnect(ctx)
	defer cancel()

	collection := client.Database(config.Database.MessageDatabase).Collection(config.Database.MessageCollection)
	message := MessageEntity{UserId: 1, RoomId: 2, Message: "hi", Time: time.Now()}

	res, err := collection.InsertOne(ctx, message)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("inserted document with ID %v\n", res.InsertedID)
}

func FindMessagesByRoomIdx(env string, roomId int) {

	client, ctx, cancel, config := mongoConn(env)
	defer client.Disconnect(ctx)
	defer cancel()

	collection := client.Database(config.Database.MessageDatabase).Collection(config.Database.MessageCollection)
	// filter := MessageEntity{RoomId: roomId}

	opts := options.Find().SetSort(bson.D{{"time", 1}})

	cursor, err := collection.Find(ctx, bson.D{{"room_id", roomId}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var results []MessageEntity
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		log.Println(result)
	}
}
