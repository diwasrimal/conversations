package routes

import (
	"log"
	"net/http"
)

func MessagesGet(w http.ResponseWriter, r *http.Request) {
	senderId := r.PathValue("sender_id")
	receiverId := r.PathValue("receiver_id")
	log.Println("sender id: ", senderId, "receiver id:", receiverId)
}
