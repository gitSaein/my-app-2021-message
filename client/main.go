package main

import (
	"flag"
	service "my-app-2021-message/service/message"
)

var env string = ""
var roomId int = 0
var userId int = 0

func init() {
	envf := flag.String("env", "local", "server environment")
	roomIdf := flag.Int("roomId", 0, "roomId")
	userIdf := flag.Int("userId", 0, "userId")
	flag.Parse()

	if flag.NFlag() == 0 { // 명령줄 옵션의 개수가 0개이면
		flag.Usage() // 명령줄 옵션 기본 사용법 출력
	}
	env = *envf
	roomId = *roomIdf
	userId = *userIdf
}

func main() {
	service.Receiver(env, roomId, userId)
}
