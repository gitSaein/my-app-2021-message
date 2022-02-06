package rabbitmq

import (
	"fmt"
	"log"
	mongoClient "my-app-2021-message/database/mongodb"
	errors "my-app-2021-message/errors"

	"github.com/rabbitmq/amqp091-go"
)

const (
	ExchangeDirect  = "direct"
	ExchangeFanout  = "fanout"
	ExchangeTopic   = "topic"
	ExchangeHeaders = "headers"
)

func Publish(mq *RabbitMQ, message string) {
	err := mq.Channel.Publish(
		"",
		mq.Queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	errors.Check(err)
}

func Consume(mq *RabbitMQ) {
	msgs, err := mq.Channel.Consume(mq.Queue.Name, "", true, false, false, false, nil)
	errors.Check(err)
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func ExchangePublish(mq *RabbitMQ, m *mongoClient.Message) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[ ERROR ]", r)
		}

	}()
	err := mq.Channel.ExchangeDeclare(
		mq.Config.Database.RabbitMQ.ExchangeName,
		ExchangeTopic,
		true,
		false,
		false,
		false,
		nil)
	errors.Check(err)

	err = mq.Channel.Publish(
		mq.Config.Database.RabbitMQ.ExchangeName,                             // exchange
		fmt.Sprintf(mq.Config.Database.RabbitMQ.MessageRoutingKey, m.RoomId), // routing key
		false, // mandatory
		false, // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(m.Message),
		})
	errors.Check(err)

	log.Printf(" [x]key [%-5s] Sent %s", fmt.Sprintf(mq.Config.Database.RabbitMQ.MessageRoutingKey, m.RoomId), m.Message)

}

func ExchangeConsume(mq *RabbitMQ, roomId int, userId int) {
	err := mq.Channel.ExchangeDeclare(
		mq.Config.Database.RabbitMQ.ExchangeName,
		ExchangeTopic,
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
	errors.Check(err)

	q, err := mq.Channel.QueueDeclare(
		fmt.Sprintf(mq.Config.Database.RabbitMQ.QueueNameByUser, userId), // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	errors.Check(err)

	log.Printf("Binding queue %s to exchange %s with routing key %s",
		q.Name, mq.Config.Database.RabbitMQ.ExchangeName,
		fmt.Sprintf(mq.Config.Database.RabbitMQ.MessageRoutingKey, roomId))
	err = mq.Channel.QueueBind(
		q.Name, // queue name
		fmt.Sprintf(mq.Config.Database.RabbitMQ.MessageRoutingKey, roomId), // routing key
		mq.Config.Database.RabbitMQ.ExchangeName,                           // exchange
		false,
		nil)
	errors.Check(err)

	msgs, err := mq.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	errors.Check(err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
