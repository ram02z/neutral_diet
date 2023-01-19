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

const deleteUserByFirebaseUID = `-- name: DeleteUserByFirebaseUID :one
DELETE FROM "user"
WHERE firebase_uid = $1
RETURNING
    id
`

func (q *Queries) DeleteUserByFirebaseUID(ctx context.Context, firebaseUid string) (int32, error) {
	row := q.db.QueryRow(ctx, deleteUserByFirebaseUID, firebaseUid)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getUserByFirebaseUID = `-- name: GetUserByFirebaseUID :one
SELECT
    id, firebase_uid, region, cf_limit, created_at, updated_at
FROM
    "user"
WHERE
    firebase_uid = $1
`

func (q *Queries) GetUserByFirebaseUID(ctx context.Context, firebaseUid string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByFirebaseUID, firebaseUid)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirebaseUid,
		&i.Region,
		&i.CfLimit,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
