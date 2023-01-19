-- name: CreateUser :one
INSERT INTO "user" (firebase_uid)
    VALUES ($1)
RETURNING
    id;

-- name: DeleteUserByFirebaseUID :one
DELETE FROM "user"
WHERE firebase_uid = $1
RETURNING
    id;

-- name: GetUserByFirebaseUID :one
SELECT
    *
FROM
    "user"
WHERE
    firebase_uid = $1;
