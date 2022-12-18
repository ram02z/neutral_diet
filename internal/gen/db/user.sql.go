// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user" (firebase_uid)
    VALUES ($1)
RETURNING
    id
`

func (q *Queries) CreateUser(ctx context.Context, firebaseUid string) (int32, error) {
	row := q.db.QueryRow(ctx, createUser, firebaseUid)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "user"
WHERE firebase_uid = $1
`

func (q *Queries) DeleteUser(ctx context.Context, firebaseUid string) error {
	_, err := q.db.Exec(ctx, deleteUser, firebaseUid)
	return err
}
