package message

import (
	"log"
	mongoClient "my-app-2021-message/database/mongodb"
	rabbitmq "my-app-2021-message/database/rabbitmq"
	errors "my-app-2021-message/errors"
)

/*
 5. 메시지 보내기
 6. 위치 정보 공유하기
*/
func Send(env string, message *mongoClient.Message) {
	client := mongoClient.Conn(env)
	defer func() {
		if r := recover(); r != nil {
			log.Println("[ ERROR ]", r)
		}
		client.Cancel()
		client.Conn.Disconnect(client.Ctx)

	}()
	errors.Check(client.Err)

	err := mongoClient.Insert(client, message)
	errors.Check(err)

	mqClient := rabbitmq.Conn(env)
	errors.Check(mqClient.Err)
	defer func() {
		mqClient.Conn.Close()
		mqClient.Channel.Close()
	}()
	rabbitmq.ExchangePublish(mqClient, message)
}

func Receiver(env string, roomId int, userId int) {
	mqClient := rabbitmq.Conn(env)
	errors.Check(mqClient.Err)

	defer func() {
		if r := recover(); r != nil {
			log.Println("[ ERROR ]", r)
		}
		mqClient.Conn.Close()
		mqClient.Channel.Close()
	}()
	rabbitmq.ExchangeConsume(mqClient, roomId, userId)
}

func GetList(env string, roomIdx int) []mongoClient.Message {
	client := mongoClient.Conn(env)
	defer func() {
		if r := recover(); r != nil {
			log.Println("[ ERROR ]", r)
		}
		client.Cancel()
		client.Conn.Disconnect(client.Ctx)

	}()
	errors.Check(client.Err)

	msgList, err := mongoClient.FindListByRoomIdx(client, roomIdx)
	errors.Check(err)

	return msgList
}
