package mongodb

import (
	"time"
)

type MessageEntity struct {
	// ID      primitive.ObjectID `bson:"_id"`
	UserId  int `bson:"user_id"`
	RoomId  int `bson:"room_id"`
	Message string
	Time    time.Time
}
