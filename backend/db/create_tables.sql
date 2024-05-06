DROP TABLE IF EXISTS users, messages, user_sessions, friends;

CREATE TABLE IF NOT EXISTS users (
	id bigserial NOT NULL PRIMARY KEY,
	fname text NOT NULL,
	lname text NOT NULL,
	username text NOT NULL UNIQUE,
	password_hash text NOT NULL,
	bio text DEFAULT ''
);

CREATE TABLE IF NOT EXISTS messages (
	id bigserial NOT NULL PRIMARY KEY,
	sender_id bigserial NOT NULL REFERENCES users(id),
	receiver_id bigserial NOT NULL REFERENCES users(id),
	text text NOT NULL,
	timestamp timestamp NOT NULL
);

-- CREATE TABLE IF NOT EXISTS sessions (
-- 	id text NOT NULL PRIMARY KEY,
-- 	user_id bigserial NOT NULL REFERENCES users(id) UNIQUE
-- );

CREATE TABLE IF NOT EXISTS user_sessions (
	user_id bigserial NOT NULL PRIMARY KEY REFERENCES users(id),
	session_id text NOT NULL
);

CREATE TABLE IF NOT EXISTS friends (
	user1_id bigserial NOT NULL REFERENCES users(id),
	user2_id bigserial NOT NULL REFERENCES users(id),
	UNIQUE(user1_id, user2_id)
);
