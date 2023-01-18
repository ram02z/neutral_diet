-- name: CreateUser :one
INSERT INTO "user" (firebase_uid)
    VALUES ($1)
RETURNING
    id;

-- name: DeleteUserByFirebaseUID :exec
DELETE FROM "user"
WHERE firebase_uid = $1;

-- name: GetUserByFirebaseUID :one
SELECT
    *
FROM
    "user"
WHERE
    firebase_uid = $1;
