package chat

import (
	"log"
	"my-app-2021-message/database/mongodb"
	"my-app-2021-message/database/rabbitmq"
	"my-app-2021-message/errors"
	"my-app-2021-message/service/message"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
1. 채팅방 생성
2. 채팅방 나가기
3. 채팅방 참여
4. 채팅방 수정 (매니저)

*/

const (
	CHAT = "chat"
)

func Create(env string, chat mongodb.Chat, participants []int) {

	mongoClient := mongodb.ApproachCollection(env, CHAT)
	errors.Check(mongoClient.Err)

	chat.Time = time.Now()

	err := mongodb.Create(mongoClient, chat)
	errors.Check(err)

	mqClient := rabbitmq.Conn(env)
	errors.Check(mqClient.Err)

	var wg sync.WaitGroup

	g := new(errgroup.Group)

	wg.Add(len(participants))
	for _, p := range participants {

		p := p
		g.Go(func() error {

			err := rabbitmq.MakeExchangeQueueBind(mqClient, p, chat.RoomId)

			return err
		})

	}
	if err := g.Wait(); err != nil {
		errors.Check(err)
	}

	message := mongodb.Message{
		UserId: chat.UserId,
		RoomId: chat.RoomId,
		Type:   message.CreateMsg,
		Time:   time.Now(),
	}

	rabbitmq.ExchangePublish(mqClient, message)
	defer func() {
		if r := recover(); r != nil {
			log.Println("[ ERROR ]", r)
		}
		mongoClient.Cancel()
		mongoClient.Conn.Disconnect(mongoClient.Ctx)
		mqClient.Conn.Close()
		mqClient.Channel.Close()

	}()
}

func Out() {

}

func In() {

}

func Update() {

}
