package db

import (
	"context"
	"log"
	"os"

	"github.com/diwasrimal/conversations/backend/models"

	"github.com/jackc/pgx/v5"
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

func CreateUser(fullname, username, passwordHash string) error {
	_, err := pool.Exec(
		context.Background(),
		"INSERT INTO users(fullname, username, password_hash) VALUES($1, $2, $3)",
		fullname,
		username,
		passwordHash,
	)
	return err
}

func UpdateUser(userId uint64, newUser models.User) error {
	_, err := pool.Exec(
		context.Background(),
		"UPDATE users SET "+
			"fullname = $1, "+
			"bio = $2 "+
			"WHERE id = $3",
		newUser.Fullname,
		newUser.Bio,
		userId,
	)
	return err
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := pool.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE username = $1",
		username,
	).Scan(&user.Id, &user.Fullname, &user.Username, &user.PasswordHash, &user.Bio); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserById(id uint64) (*models.User, error) {
	var user models.User
	if err := pool.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE id = $1",
		id,
	).Scan(&user.Id, &user.Fullname, &user.Username, &user.PasswordHash, &user.Bio); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserBySessionId(sessionId string) (*models.User, error) {
	var user models.User
	if err := pool.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE id = ( "+
			"SELECT user_id FROM user_sessions WHERE session_id = $1 "+
			")",
		sessionId,
	).Scan(&user.Id, &user.Fullname, &user.Username, &user.PasswordHash, &user.Bio); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
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
	var session models.Session
	if err := pool.QueryRow(
		context.Background(),
		"SELECT * FROM user_sessions WHERE session_id = $1",
		sessionId,
	).Scan(&session.UserId, &session.SessionId); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &session, nil
}

func GetMessagesAmong(userId1, userId2 uint64) ([]models.Message, error) {
	var messages []models.Message
	rows, err := pool.Query(
		context.Background(),
		"SELECT * FROM messages WHERE "+
			"(sender_id = $1 AND receiver_id = $2) OR "+
			"(sender_id = $2 AND receiver_id = $1)"+
			"ORDER BY timestamp DESC",
		userId1,
		userId2,
	)
	if err != nil {
		return messages, err
	}
	defer rows.Close()
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.Id, &msg.SenderId, &msg.ReceiverId, &msg.Text, &msg.Timestamp); err != nil {
			return messages, err
		}
		messages = append(messages, msg)
	}
	return messages, nil // TODO: maybe add limit
}

func GetConversationsOf(userId uint64) ([]models.Conversation, error) {
	var conversations []models.Conversation
	rows, err := pool.Query(
		context.Background(),
		"SELECT * FROM conversations WHERE "+
			"user1_id = $1 OR user2_id = $1 "+
			"ORDER BY timestamp DESC",
		userId,
	)
	if err != nil {
		return conversations, err
	}
	defer rows.Close()
	for rows.Next() {
		var conv models.Conversation
		if err := rows.Scan(&conv.UserId1, &conv.UserId2, &conv.Timestamp); err != nil {
			return conversations, err
		}
		conversations = append(conversations, conv)
	}
	return conversations, nil // TODO: maybe add limit
}

func GetRecentChatPartners(userId uint64) ([]models.User, error) {
	var partners []models.User
	rows, err := pool.Query(
		context.Background(),
		`SELECT * FROM users WHERE id IN (
			SELECT CASE WHEN user1_id = $1 THEN user2_id ELSE user1_id END
			FROM conversations WHERE
			user1_id = $1 OR user2_id = $1
			ORDER BY timestamp DESC
		)`,
		userId,
	)
	if err != nil {
		return partners, nil
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Fullname, &user.Username, &user.PasswordHash, &user.Bio); err != nil {
			return partners, err
		}
		partners = append(partners, user)
	}
	return partners, nil // TODO: maybe add limit
}

func SearchUser(searchType, searchQuery string) ([]models.User, error) {
	var matches []models.User
	var rows pgx.Rows
	var err error

	// Fuzzy search using likeness and levenshtein distance.
	if searchType == "normal" {
		maxLevDist := int(0.5 * float64(len(searchQuery)))
		log.Println("maxLevList:", maxLevDist)
		rows, err = pool.Query(
			context.Background(),
			` SELECT * FROM users WHERE
				fullname ILIKE '%' || $1 || '%' OR
				levenshtein(fullname, $1) <= $2
				ORDER BY levenshtein(fullname, $1) ASC;
			`,
			searchQuery,
			maxLevDist,
		)
	} else if searchType == "by-username" {
		rows, err = pool.Query(
			context.Background(),
			"SELECT * FROM users WHERE username ILIKE '%' || $1 || '%'",
			searchQuery,
		)
	}
	if err != nil {
		return matches, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Fullname, &user.Username, &user.PasswordHash, &user.Bio); err != nil {
			return matches, err
		}
		matches = append(matches, user)
	}
	return matches, nil // TODO: maybe add limit
}

func RecordFriendRequest(from, to uint64) error {
	_, err := pool.Exec(
		context.Background(),
		`INSERT INTO friend_requests(requestor_id, receiver_id)
			VALUES($1, $2)
			ON CONFLICT DO NOTHING`,
		from,
		to,
	)
	return err
}

func DeleteFriendRequest(from, to uint64) error {
	_, err := pool.Exec(
		context.Background(),
		`DELETE FROM friend_requests WHERE
			requestor_id = $1 AND receiver_id = $2`,
		from,
		to,
	)
	return err
}

func RecordFriendship(userId1, userId2 uint64) error {
	_, err := pool.Exec(
		context.Background(),
		"INSERT INTO friends(user1_id, user2_id) VALUES($1, $2)",
		userId1,
		userId2,
	)
	return err
}

func DeleteFriendship(userId1, userId2 uint64) error {
	_, err := pool.Exec(
		context.Background(),
		`DELETE FROM friends WHERE
			user1_id = $1 AND user2_id = $2 OR
			user1_id = $2 AND user2_id = $1`,
		userId1,
		userId2,
	)
	return err
}

// Returns status of friendship for two users from first user's point of view.
// Can give 4 statuses, "friends", "req-sent", "req-received", "unknown".
// Ex. "req-sent" means first user has sent a request to second.
func GetFriendshipStatus(userId, otherUserId uint64) (string, error) {
	var status string
	if err := pool.QueryRow(
		context.Background(),
		`SELECT CASE
			WHEN EXISTS (
				SELECT 1 FROM friends WHERE
				(user1_id = $1 AND user2_id = $2) OR
				(user2_id = $1 AND user1_id = $2) ) THEN 'friends'
			WHEN EXISTS (
				SELECT 1 FROM friend_requests WHERE requestor_id = $1 AND receiver_id = $2
			) THEN 'req-sent'
			WHEN EXISTS (
				SELECT 1 FROM friend_requests WHERE receiver_id = $1 AND requestor_id = $2
			) THEN 'req-received'
			ELSE 'unknown'
		END AS status`,
		userId,
		otherUserId,
	).Scan(&status); err != nil {
		return "", err
	}
	return status, nil
}

// Returns list of users that are friends to user with given id
func GetFriends(userId uint64) ([]models.User, error) {
	var friends []models.User
	rows, err := pool.Query(
		context.Background(),
		`SELECT * FROM users WHERE id IN (
			SELECT CASE WHEN user1_id = $1 THEN user2_id ELSE user1_id END
			FROM friends WHERE
			user1_id = $1 OR user2_id = $1
		)`,
		userId,
	)
	if err != nil {
		return friends, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Fullname, &user.Username, &user.PasswordHash, &user.Bio); err != nil {
			return friends, err
		}
		friends = append(friends, user)
	}
	return friends, nil // TODO: maybe add limit
}

// Returns list of users that have sent request to given user
func GetFriendRequestorsTo(userId uint64) ([]models.User, error) {
	var requestors []models.User
	rows, err := pool.Query(
		context.Background(),
		`SELECT * FROM users WHERE id IN (
			SELECT requestor_id FROM friend_requests WHERE
			receiver_id = $1
		)`,
		userId,
	)
	if err != nil {
		return requestors, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Fullname, &user.Username, &user.PasswordHash, &user.Bio); err != nil {
			return requestors, err
		}
		requestors = append(requestors, user)
	}
	return requestors, nil // TODO: maybe add limit
}
