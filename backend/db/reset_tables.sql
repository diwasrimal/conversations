TRUNCATE users, messages, user_sessions, friends RESTART IDENTITY;

-- Fill with dummy data
INSERT INTO users(fname, lname, username, password_hash) VALUES
	('Ram', 'Poudel', 'ram', '$2a$10$vT5JwJLYO6O7Tg5sIMeca.6MIKQGpybEwiLwWvo4n5Ae.TUMeG73y'), -- ram (bcrypted with default cost)
	('Shyam', 'Acharya', 'shyam', '$2a$10$kvRD/jewfMPNt3oMKMGPpe9Hz6pTbWOI1GTfEqT9KsZw160WR5TZO'), -- shyam
	('Hari', 'Rai', 'hari', '$2a$10$U32d4KQn4nR.vBEcHKIwt.WS3DGU16SRlNUSg4pDO3iqAzX23aKfO'), -- hari
	('Mohan', 'Devkota', 'mohan', '$2a$10$bonHqtMXx4V6q550CE4AwuhPx5kF2RPCf7QqJZLSdd9aguVfbn4Ta'), -- mohan
	('Rita', 'Gurung', 'rita', '$2a$10$v49I8/eiUOf.u6jflGiFw.zSlIB0NLbJFB8yB7PZbAUllxPc737bC'); -- rita

INSERT INTO friends VALUES
	(1, 2),
	(2, 1),
	(1, 3),
	(3, 1),
	(4, 5),
	(5, 4);

INSERT INTO messages(sender_id, receiver_id, text, timestamp) VALUES
	(1, 2, 'user 1 to 2: hello user 2', 'Sun Jan  5 20:02:53 +0545 2024'),
	(2, 1, 'user 2 to 1: hi user 1', 'Sun May  3 20:02:53 +0545 2024'),
	(2, 1, 'user 2 to 1: how are you doing?', 'Sun May  4 20:02:53 +0545 2024');
	