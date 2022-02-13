package main

import (
	"flag"
	"fmt"
	"my-app-2021-message/api"
	"net/http"

	"github.com/gorilla/mux"
)

var Env string = ""

func init() {
	envf := flag.String("env", "local", "server environment")
	flag.Parse()

	if flag.NFlag() == 0 { // 명령줄 옵션의 개수가 0개이면
		flag.Usage() // 명령줄 옵션 기본 사용법 출력
	}
	Env = *envf
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/messages", api.SendMessage(Env)).Methods("POST")
	router.HandleFunc("/messages/{roomIdx}", api.GetMessages(Env)).Methods("GET")
	router.HandleFunc("/chat/{type}", api.Chat(Env)).Methods("POST")

	http.ListenAndServe(":8084", router)

}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
