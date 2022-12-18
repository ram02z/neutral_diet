-- name: CreateUser :one
INSERT INTO "user" (firebase_uid)
    VALUES ($1)
RETURNING
    id;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE firebase_uid = $1;
