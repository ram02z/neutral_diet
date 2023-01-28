-- name: CreateUser :one
INSERT INTO "user" (firebase_uid, region, cf_limit)
    VALUES ($1, $2, $3)
RETURNING
    id;

-- name: DeleteUserByFirebaseUID :one
DELETE FROM "user"
WHERE firebase_uid = $1
RETURNING
    id;

-- name: UpdateUserSettings :exec
UPDATE
    "user"
SET
    region = $2,
    cf_limit = $3
WHERE
    firebase_uid = $1;

-- name: GetUserByFirebaseUID :one
SELECT
    *
FROM
    "user"
WHERE
    firebase_uid = $1;
