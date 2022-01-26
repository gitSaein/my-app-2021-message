package main

import (
	"flag"
	mongoClient "my-app-2021-message/database/mongodb"
	rabbitmq "my-app-2021-message/database/rabbitmq"
)

var env string = ""

func init() {
	envf := flag.String("env", "local", "server environment")
	flag.Parse()

	if flag.NFlag() == 0 { // 명령줄 옵션의 개수가 0개이면
		flag.Usage() // 명령줄 옵션 기본 사용법 출력
	}
	env = *envf
}

func main() {

	client := mongoClient.MongoConn(env)
	defer client.Conn.Disconnect(client.Ctx)
	defer client.Cancel()
	if client.Err != nil {
		panic(client.Err)
	}

	mongoClient.InsertMessage(client)
	mongoClient.FindMessagesByRoomIdx(client, 2)
	rabbitmq.ConnRabbitMq(env)

}
