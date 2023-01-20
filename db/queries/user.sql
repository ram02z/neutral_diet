-- name: CreateUser :one
INSERT INTO "user" (firebase_uid, region)
    VALUES ($1, $2)
RETURNING
    id;

-- name: DeleteUserByFirebaseUID :one
DELETE FROM "user"
WHERE firebase_uid = $1
RETURNING
    id;

-- name: UpdateUserRegion :exec
UPDATE
    "user"
SET
    region = $2
WHERE
    firebase_uid = $1;

-- name: GetUserByFirebaseUID :one
SELECT
    *
FROM
    "user"
WHERE
    firebase_uid = $1;
