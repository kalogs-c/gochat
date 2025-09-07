-- name: GetRoom :one
SELECT * FROM rooms
WHERE id = ?;

-- name: CreateRoom :one
INSERT INTO rooms (topic) VALUES (?)
RETURNING *;
