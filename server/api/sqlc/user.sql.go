// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    hashed_password
) VALUES (
  $1, $2, $3
) RETURNING id, name, email, hashed_password, password_changed_at, updated_at, created_at
`

type CreateUserParams struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, hashed_password, password_changed_at, updated_at, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, name, email, hashed_password, password_changed_at, updated_at, created_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUserParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUser(ctx context.Context, arg ListUserParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUser, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.HashedPassword,
			&i.PasswordChangedAt,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users 
SET name = $2,email = $3
WHERE id = $1
RETURNING id, name, email, hashed_password, password_changed_at, updated_at, created_at
`

type UpdateUserParams struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ID, arg.Name, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
