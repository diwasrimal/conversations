package routes

import (
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/diwasrimal/conversations/backend/db"
	"github.com/diwasrimal/conversations/backend/models"
	"github.com/diwasrimal/conversations/backend/types"
	"github.com/gorilla/websocket"
)

// Stores the connection of each client
var clientsMu sync.RWMutex
var clients = make(map[uint64]*websocket.Conn)

var wsUp = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WSHandleFunc(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(uint64)
	log.Printf("Hit WSHandleFunc() with userId: %v\n", userId)
	conn, err := wsUp.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to ws: %v\n", err)
		return
	}

	clientsMu.Lock()
	clients[userId] = conn
	clientsMu.Unlock()

	// Remove the connection from map and close
	defer func() {
		clientsMu.Lock()
		delete(clients, userId)
		clientsMu.Unlock()
		conn.Close()
	}()

	for true {
		var data types.Json
		err := conn.ReadJSON(&data)
		if err != nil {
			log.Printf("%T reading ws json data: %v\n", err, err)
			break
		}

		log.Printf("ws data: %v\n", data)

		switch data["msgType"] {
		case "chatMessageSend":
			if err := handleMsgSend(userId, data); err != nil {
				log.Printf("Error handling msg send: %v\n", err)
			}
		}
	}
}

func handleMsgSend(senderId uint64, data types.Json) error {
	rid, ridOk := data["receiverId"].(float64)
	text, textOk := data["text"].(string)
	ts, tsOk := data["timestamp"].(string)
	if !ridOk || !textOk || !tsOk {
		return errors.New("Invalid/Missing data fields, ")
	}

	receiverId := uint64(rid)
	timestamp, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return err
	}

	// Broadcast message to other user if active
	clientsMu.RLock()
	peerConn, ok := clients[receiverId]
	clientsMu.RUnlock()
	if ok {
		peerConn.WriteJSON(types.Json{
			"msgType":   "chatMessageReceive",
			"senderId":  senderId,
			"text":      text,
			"timestamp": timestamp,
		})
	}

	// Store the message in database
	if err := db.RecordMessage(models.Message{
		SenderId:   senderId,
		ReceiverId: receiverId,
		Text:       text,
		Timestamp:  timestamp,
	}); err != nil {
		log.Printf("Error recording message of %v in db: %v\n", senderId, err)
	}

	// Update the last conversation time for two users
	if err := db.UpdateOrCreateConversation(senderId, receiverId, timestamp); err != nil {
		log.Printf("Error updating last conv for (%v,%v): %v\n", senderId, receiverId, err)
	}

	return nil
}
