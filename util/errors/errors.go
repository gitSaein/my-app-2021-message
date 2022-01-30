package errors

import (
	"encoding/json"
	"net/http"
	"time"
)

type Errors struct {
	Errors []*Error `json:"errors"`
}
type Error struct {
	Id          string    `json:"id"`
	Status      int       `json:"status"`
	Detail      string    `json:"detail"`
	OriginError error     `json:"originError"`
	Timestamp   time.Time `json:"timestamp"`
}

func WriteError(w http.ResponseWriter, err *Error) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(Errors{[]*Error{err}})
}

var (
	ErrBadRequest     = &Error{Id: "bad_request", Status: 400, Detail: "Request body is not well-formed. It must be JSON."}
	ErrInternalServer = &Error{Id: "internal_server_error", Status: 500, Timestamp: time.Now()}
)

func Check(err error) bool {
	switch err.(type) {
	case error:
		ErrInternalServer.OriginError = err
		ErrInternalServer.Detail = err.Error()
		panic(ErrInternalServer)
	default:
		return true
	}
}
