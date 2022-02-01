package mongodb

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	log.Println("start mongo")
}

type Message struct {
	ID      interface{} `bson:"_id,omitempty"`
	UserId  int         `bson:"user_id"`
	RoomId  int         `bson:"room_id"`
	Message string
	Time    time.Time
}

func Insert(connInfo *ConnInfo, message *Message) error {
	message.Time = time.Now()
	res, err := connInfo.Collection.InsertOne(connInfo.Ctx, message)
	if err != nil {
		return err
	}
	message.ID = res.InsertedID
	return nil
}

func FindListByRoomIdx(connInfo *ConnInfo, roomId int) ([]Message, error) {
	var results []Message

	opts := options.Find().SetSort(bson.D{{"time", 1}})

	cursor, err := connInfo.Collection.Find(connInfo.Ctx,
		bson.D{{"room_id", roomId}}, opts)
	if err != nil {
		return results, err
	}
	defer cursor.Close(connInfo.Ctx)

	if err = cursor.All(connInfo.Ctx, &results); err != nil {
		return results, err
	}
	for _, result := range results {
		log.Println(result)
	}
	return results, nil
}
