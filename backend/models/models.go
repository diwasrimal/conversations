package models

import "time"

type User struct {
	Id           uint64 `json:"id"`
	Fname        string `json:"fname"`
	Lname        string `json:"lname"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
	Bio          string `json:"bio"`
}

type Session struct {
	UserId    uint64
	SessionId string
}

type Message struct {
	Id         uint64    `json:"-"`
	SenderId   uint64    `json:"senderId"`
	ReceiverId uint64    `json:"receiverId"`
	Text       string    `json:"text"`
	Timestamp  time.Time `json:"timestamp"`
}
