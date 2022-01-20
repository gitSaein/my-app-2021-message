package mongodb

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageEntity struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserId  int                `bson:"user_id"`
	RoomId  int                `bson:"room_id"`
	Message string
	Time    time.Time
}
