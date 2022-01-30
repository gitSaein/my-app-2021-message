package service

import (
	"log"
	mongoClient "my-app-2021-message/database/mongodb"
	rabbitmq "my-app-2021-message/database/rabbitmq"
	errors "my-app-2021-message/util/errors"
)

/*
 5. 메시지 보내기
 6. 위치 정보 공유하기
*/
func Send(env string, message mongoClient.MessageEntity) {
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
	rabbitmq.Publish(mqClient, message.Message)
}

func Receiver(env string) {
	mqClient := rabbitmq.Conn(env)
	defer func() {
		mqClient.Conn.Close()
		mqClient.Channel.Close()
	}()
	rabbitmq.Consume(mqClient)
}
