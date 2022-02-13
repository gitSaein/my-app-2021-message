package mongodb

import (
	"log"
	"time"
)

func init() {
	log.Println("start mongo")
}

type Chat struct {
	ID       interface{} `bson:"_id,omitempty"`
	RoomId   int         `bson:"room_id"`
	RoomName string      `bson:"rom_name"`
	Time     time.Time
}

type Participants struct {
	ID         interface{} `bson:"_id,omitempty"`
	RoomId     int         `bson:"room_id"`
	UserId     string      `bson:"user_id"`
	CreateDate time.Time   `bson:"create_date"`
	UpdateDate time.Time   `bson:"update_date"`
}

func Create(connInfo *ConnInfo, chat Chat) error {
	chat.Time = time.Now()
	res, err := connInfo.Collection.InsertOne(connInfo.Ctx, chat)
	if err != nil {
		return err
	}
	chat.ID = res.InsertedID
	return nil
}
