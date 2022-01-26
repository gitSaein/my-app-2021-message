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

func descSort() {

}

func InsertMessage(connInfo *ConnInfo) {

	message := MessageEntity{UserId: 1, RoomId: 2, Message: "hi", Time: time.Now()}

	res, err := connInfo.Collection.InsertOne(connInfo.Ctx, message)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("inserted document with ID %v\n", res.InsertedID)
}

func FindMessagesByRoomIdx(connInfo *ConnInfo, roomId int) ([]MessageEntity, error) {
	var results []MessageEntity

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
