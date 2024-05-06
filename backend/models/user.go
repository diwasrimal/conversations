package models

type User struct {
	Id           uint64
	Fname        string
	Lname        string
	Username     string
	PasswordHash string
	Bio          string
}
