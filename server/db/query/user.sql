-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    hashed_password
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users 
SET name = $2,email = $3
WHERE id = $1
RETURNING *;

-- name: ListUser :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;
