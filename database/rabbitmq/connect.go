package rabbitmq

import (
	"log"
	conf "my-app-2021-message/config"
	errors "my-app-2021-message/util/errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Channel *amqp.Channel
	Conn    *amqp.Connection
	Queue   amqp.Queue
}

func Conn(env string) *RabbitMQ {
	log.Println("rabbitmq Conn start...")

	conf := conf.GetCongif(env)
	conn, err := amqp.Dial(conf.Database.RabbitMQ.Url)
	errors.Check(err)

	ch, err := conn.Channel()
	errors.Check(err)

	q, err := ch.QueueDeclare("testQueue", false, false, false, false, nil)
	errors.Check(err)

	return &RabbitMQ{Channel: ch, Queue: q, Conn: conn}

}
