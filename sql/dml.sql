INSERT INTO user (name, email, nick, password) VALUES
('User 1', 'user1@email.com', 'user1', '$2a$10$YbckPh9p13qTF7KzcfllV.kFpZ7okeqIbIyVfjvrcoMmQZbNJSIxy'), # Password is 123 for all users in dml.sql
('User 2', 'user2@email.com', 'user2', '$2a$10$YbckPh9p13qTF7KzcfllV.kFpZ7okeqIbIyVfjvrcoMmQZbNJSIxy'),
('User 3', 'user3@email.com', 'user3', '$2a$10$YbckPh9p13qTF7KzcfllV.kFpZ7okeqIbIyVfjvrcoMmQZbNJSIxy');

INSERT INTO follower (user_id, follower_id) VALUES 
(1, 2),
(2, 1), 
(3, 1);