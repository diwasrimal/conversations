package models

import "time"

type User struct {
	Id           uint64
	Fname        string
	Lname        string
	Username     string
	PasswordHash string
	Bio          string
}

type Session struct {
	UserId    uint64
	SessionId string
}

type Message struct {
	Id         uint64
	SenderId   uint64
	ReceiverId uint64
	Text       string
	Timestamp  time.Time
}
