package api

import (
	"my-app-2021-message/database/mongodb"
)

type RequestChat struct {
	Chat         mongodb.Chat
	Participants []int
}
