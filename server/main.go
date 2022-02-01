package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	mongodb "my-app-2021-message/database/mongodb"
	"my-app-2021-message/errors"
	msg "my-app-2021-message/service/message"
	"net/http"
	"strconv"

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

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func SendMessage(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			respondWithJSON(w, http.StatusInternalServerError, r)
			log.Println("[ ERROR ]", r)

		}
	}()

	var message mongodb.Message

	err := json.NewDecoder(r.Body).Decode(&message)
	errors.Check(err)
	msg.Send(Env, &message)
	respondWithJSON(w, http.StatusOK, &message)

}

func GetMessages(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			respondWithJSON(w, http.StatusInternalServerError, r)
			log.Println("[ ERROR ]", r)

		}
	}()

	vars := mux.Vars(r)
	idx, err := strconv.Atoi(vars["roomIdx"])
	errors.Check(err)
	messages := msg.GetList(Env, idx)
	respondWithJSON(w, http.StatusOK, &messages)

}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/messages", SendMessage).Methods("POST")
	router.HandleFunc("/messages/{roomIdx}", GetMessages).Methods("GET")

	http.ListenAndServe(":8084", router)

}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
