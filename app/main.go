package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func sendMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Message: %v\n", vars)
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/messages/send", sendMessage)
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
