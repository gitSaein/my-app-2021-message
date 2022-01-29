package rabbitmq

import (
	"log"
	errors "my-app-2021-message/util/errors"

	"github.com/rabbitmq/amqp091-go"
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
