package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Username     string
	PasswordHash string
}

type Session struct {
	Id       string
	Username string
}

var pool *pgxpool.Pool

func MustInit() {
	url, set := os.LookupEnv("DATABASE_URL")
	if !set {
		panic("environment variable 'DATABASE_URL' not set")
	}
	var err error
	pool, err = pgxpool.New(context.Background(), url)
	if err != nil {
		panic(err)
	}

	_, err = pool.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS users (
			username text NOT NULL PRIMARY KEY,
			password_hash text NOT NULL
		)`,
	)
	if err != nil {
		panic(err)
	}
	_, err = pool.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS sessions (
			id text NOT NULL PRIMARY KEY,
			username text NOT NULL REFERENCES users (username)
		)`,
	)
	if err != nil {
		panic(err)
	}
}

func Close() {
	pool.Close()
}

func LookupUser(username string) (*User, error) {
	rows, err := pool.Query(
		context.Background(),
		`SELECT * FROM users WHERE username = $1`,
		username,
	)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}
	values, err := rows.Values()
	if err != nil {
		return nil, err
	}
	return &User{
		Username:     values[0].(string),
		PasswordHash: values[1].(string),
	}, nil
}

func RecordUser(user *User) error {
	_, err := pool.Exec(
		context.Background(),
		`INSERT INTO users(username, password_hash) VALUES($1, $2)`,
		user.Username,
		user.PasswordHash,
	)
	return err
}

func IsUsernameTaken(username string) (bool, error) {
	user, err := LookupUser(username)
	if err != nil {
		return false, err
	}
	userExists := user != nil
	return userExists, nil
}

func RecordSession(id string, username string) error {
	_, err := pool.Exec(
		context.Background(),
		`INSERT INTO sessions(id, username) VALUES($1, $2)`,
		id,
		username,
	)
	return err
}

func LookupSession(id string) (*Session, error) {
	rows, err := pool.Query(
		context.Background(),
		`SELECT * FROM sessions WHERE id = $1`,
		id,
	)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}
	values, err := rows.Values()
	if err != nil {
		return nil, err
	}
	return &Session{
		Id:       values[0].(string),
		Username: values[1].(string),
	}, nil
}
