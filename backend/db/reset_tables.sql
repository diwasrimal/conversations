TRUNCATE users, messages, user_sessions, conversations, friends, friend_requests RESTART IDENTITY;

-- Fill with dummy data
INSERT INTO users(fullname, username, password_hash) VALUES
	('Ram Poudel', 'ram', '$2a$10$vT5JwJLYO6O7Tg5sIMeca.6MIKQGpybEwiLwWvo4n5Ae.TUMeG73y'), -- ram (bcrypted with default cost)
	('Shyam Acharya', 'shyam', '$2a$10$kvRD/jewfMPNt3oMKMGPpe9Hz6pTbWOI1GTfEqT9KsZw160WR5TZO'), -- shyam
	('Hari Rai', 'hari', '$2a$10$U32d4KQn4nR.vBEcHKIwt.WS3DGU16SRlNUSg4pDO3iqAzX23aKfO'), -- hari
	('Mohan Devkota', 'mohan', '$2a$10$bonHqtMXx4V6q550CE4AwuhPx5kF2RPCf7QqJZLSdd9aguVfbn4Ta'), -- mohan
	('Rita Gurung', 'rita', '$2a$10$v49I8/eiUOf.u6jflGiFw.zSlIB0NLbJFB8yB7PZbAUllxPc737bC'); -- rita

INSERT INTO friends VALUES
	(1, 2),
	(1, 3),
	(4, 5);

INSERT INTO messages(sender_id, receiver_id, text, timestamp) VALUES
	(1, 2, 'user 1 to 2: hello user 2', '2024-01-05 20:02:53'),
	(2, 1, 'user 2 to 1: hi user 1', '2024-05-03 20:02:53'),
	(2, 1, 'user 2 to 1: how are you doing?', '2024-05-04 20:02:53'),
	(1, 3, 'user 1 to 3: hello user 3', '2024-10-01 01:01:53'),
	(2, 4, 'user 2 to 4: hello user 4', '2023-10-01 01:01:53');

INSERT INTO conversations(user1_id, user2_id, timestamp) VALUES
	(1, 2, '2024-05-04 20:02:53'), -- latest
	(1, 3, '2024-10-01 01:01:53'),
	(2, 4, '2023-10-01 01:01:53');