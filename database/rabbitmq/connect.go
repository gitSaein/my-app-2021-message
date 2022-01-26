package rabbitmq

import (
	"fmt"
	"log"
	conf "my-app-2021-message/config"
	errors "my-app-2021-message/util/errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnRabbitMq(env string) {
	conf := conf.GetCongif(env)
	conn, err := amqp.Dial(conf.Database.RabbitMQ.Url)
	errors.Check(err)
	if err != nil {
		log.Fatalf("Failed Initializing Broker Connection: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	errors.Check(err)
	defer ch.Close()

	q, err := ch.QueueDeclare("testQueue", false, false, false, false, nil)
	errors.Check(err)

	fmt.Println(q)

	err = ch.Publish("", "testQueue", false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("hello"),
		})
	errors.Check(err)

}
