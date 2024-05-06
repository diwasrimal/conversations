package db

import (
	"context"
	"log"
	"os"

	"github.com/diwasrimal/gochat/backend/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func MustInit() {
	url, set := os.LookupEnv("DATABASE_URL")
	if !set {
		panic("Environment variable 'DATABASE_URL' not set!")
	}
	var err error
	pool, err = pgxpool.New(context.Background(), url)
	if err != nil {
		panic(err)
	}
	log.Printf("Intialized db %q", url)
}

func Close() {
	pool.Close()
	log.Println("Closed db")
}

func CreateUser(fname, lname, username, passwordHash string) error {
	_, err := pool.Exec(
		context.Background(),
		"INSERT INTO users(fname, lname, username, password_hash) VALUES($1, $2, $3, $4)",
		fname,
		lname,
		username,
		passwordHash,
	)
	return err
}

func UpdateUser(userId uint64, newUser models.User) error {
	_, err := pool.Exec(
		context.Background(),
		"UPDATE users SET "+
			"fname = $1, "+
			"lname = $2, "+
			"bio = $3 "+
			"WHERE id = $4",
		newUser.Fname,
		newUser.Lname,
		newUser.Bio,
		userId,
	)
	return err
}

func GetUserByUsername(username string) (*models.User, error) {
	rows, err := pool.Query(
		context.Background(),
		"SELECT * FROM users WHERE username = $1",
		username,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	user := models.User{}
	err = rows.Scan(&user.Id, &user.Fname, &user.Lname, &user.Username, &user.PasswordHash, &user.Bio)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(id uint64) (*models.User, error) {
	rows, err := pool.Query(
		context.Background(),
		"SELECT * FROM users WHERE id = $1",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	user := models.User{}
	err = rows.Scan(&user.Id, &user.Fname, &user.Lname, &user.Username, &user.PasswordHash, &user.Bio)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserBySessionId(sessionId string) (*models.User, error) {
	rows, err := pool.Query(
		context.Background(),
		"SELECT * FROM users WHERE id = ( "+
			"SELECT user_id FROM user_sessions WHERE session_id = $1 "+
			")",
		sessionId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	user := models.User{}
	err = rows.Scan(&user.Id, &user.Fname, &user.Lname, &user.Username, &user.PasswordHash, &user.Bio)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func IsUsernameTaken(username string) (bool, error) {
	rows, err := pool.Query(
		context.Background(),
		"SELECT username FROM users WHERE username = $1",
		username,
	)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}

func CreateUserSession(userId uint64, sessionId string) error {
	_, err := pool.Exec(
		context.Background(),
		"INSERT INTO user_sessions(user_id, session_id) "+
			"VALUES($1, $2) "+
			"ON CONFLICT(user_id) DO UPDATE "+
			"SET session_id = excluded.session_id",
		userId,
		sessionId,
	)
	return err
}

func DeleteUserSession(sessionId string) error {
	_, err := pool.Exec(
		context.Background(),
		"DELETE FROM user_sessions WHERE session_id = $1",
		sessionId,
	)
	return err
}

func GetSession(sessionId string) (*models.Session, error) {
	rows, err := pool.Query(
		context.Background(),
		"SELECT * FROM user_sessions WHERE session_id = $1",
		sessionId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	session := models.Session{}
	err = rows.Scan(&session.UserId, &session.SessionId)
	if err != nil {
		return nil, err
	}
	return &session, nil
}