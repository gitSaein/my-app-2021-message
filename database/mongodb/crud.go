package mongodb

import (
	"context"
	"log"
	"time"
)

func Insert() {
	// var coll *mongo.Collection
	client := mongoConn()
	defer client.Disconnect(context.TODO())

	collection := client.Database("admin").Collection("message")
	message := MessageEntity{UserId: 1, RoomId: 2, Message: "hi", Time: time.Now()}

	res, err := collection.InsertOne(context.TODO(), message)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("inserted document with ID %v\n", res.InsertedID)
}
