package api

import (
	"encoding/json"
	"log"
	"my-app-2021-message/errors"
	c "my-app-2021-message/service/chat"
	msg "my-app-2021-message/service/message"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	CreateChat = "create"
	InChat     = "in"
	OutChat    = "out"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func SendMessage(env string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := msg.SendMsg(env, r.Body)
		respondWithJSON(w, http.StatusOK, &message)
	}
}

func GetMessages(env string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				respondWithJSON(w, http.StatusBadRequest, r)
				log.Printf("[ERROR][%d] %v", http.StatusBadRequest, r)

			}
		}()

		vars := mux.Vars(r)
		idx, err := strconv.Atoi(vars["roomIdx"])
		errors.Check(err)
		messages := msg.GetList(env, idx)
		respondWithJSON(w, http.StatusOK, &messages)
	}
}

func Chat(env string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				respondWithJSON(w, http.StatusBadRequest, r)
				log.Printf("[ERROR][%d] %v", http.StatusBadRequest, r)

			}
		}()

		vars := mux.Vars(r)
		var reqChat RequestChat

		err := json.NewDecoder(r.Body).Decode(&reqChat)
		errors.Check(err)

		switch vars["type"] {
		case CreateChat:
			c.Create(env, reqChat.Chat, reqChat.Participants)
			respondWithJSON(w, http.StatusOK, &reqChat)

		case InChat:
			c.In()
			respondWithJSON(w, http.StatusOK, &reqChat)

		case OutChat:
			c.Out()
			respondWithJSON(w, http.StatusOK, &reqChat)

		default:
			respondWithJSON(w, http.StatusBadRequest, r)
			log.Printf("[ERROR][%d] %v", http.StatusBadRequest, r)
		}
	}
}
