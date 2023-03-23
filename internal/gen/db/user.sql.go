// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: user.sql

package db

import (
	"context"

	"github.com/shopspring/decimal"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user" (firebase_uid, region, cf_limit, dietary_requirement)
    VALUES ($1, $2, $3, $4)
RETURNING
    id
`

type CreateUserParams struct {
	FirebaseUid        string
	Region             int32
	CfLimit            decimal.Decimal
	DietaryRequirement int32
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.FirebaseUid,
		arg.Region,
		arg.CfLimit,
		arg.DietaryRequirement,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteUserByFirebaseUID = `-- name: DeleteUserByFirebaseUID :exec
DELETE FROM "user"
WHERE firebase_uid = $1
`

func (q *Queries) DeleteUserByFirebaseUID(ctx context.Context, firebaseUid string) error {
	_, err := q.db.Exec(ctx, deleteUserByFirebaseUID, firebaseUid)
	return err
}

const getUserByFirebaseUID = `-- name: GetUserByFirebaseUID :one
SELECT
    id, firebase_uid, region, cf_limit, created_at, updated_at, dietary_requirement
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
		&i.DietaryRequirement,
	)
	return i, err
}

const updateUserSettings = `-- name: UpdateUserSettings :exec
UPDATE
    "user"
SET
    region = $2,
    cf_limit = $3,
    dietary_requirement = $4
WHERE
    firebase_uid = $1
`

type UpdateUserSettingsParams struct {
	FirebaseUid        string
	Region             int32
	CfLimit            decimal.Decimal
	DietaryRequirement int32
}

func (q *Queries) UpdateUserSettings(ctx context.Context, arg UpdateUserSettingsParams) error {
	_, err := q.db.Exec(ctx, updateUserSettings,
		arg.FirebaseUid,
		arg.Region,
		arg.CfLimit,
		arg.DietaryRequirement,
	)
	return err
}
