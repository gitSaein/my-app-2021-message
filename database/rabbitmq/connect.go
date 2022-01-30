package rabbitmq

import (
	"log"
	conf "my-app-2021-message/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Channel *amqp.Channel
	Conn    *amqp.Connection
	Queue   amqp.Queue
	Err     error
}

func Conn(env string) *RabbitMQ {
	log.Println("rabbitmq Conn start...")

	conf := conf.GetCongif(env)
	conn, err := amqp.Dial(conf.Database.RabbitMQ.Url)
	if err != nil {
		return &RabbitMQ{Err: err}

	}

	ch, err := conn.Channel()
	if err != nil {
		return &RabbitMQ{Err: err}

	}
	q, err := ch.QueueDeclare(conf.Database.RabbitMQ.QueueName, false, false, false, false, nil)
	if err != nil {
		return &RabbitMQ{Err: err}

	}
	return &RabbitMQ{Channel: ch, Queue: q, Conn: conn}

}
