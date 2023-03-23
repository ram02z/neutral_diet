-- name: CreateUser :one
INSERT INTO "user" (firebase_uid, region, cf_limit, dietary_requirement)
    VALUES ($1, $2, $3, $4)
RETURNING
    id;

-- name: DeleteUserByFirebaseUID :exec
DELETE FROM "user"
WHERE firebase_uid = $1;

-- name: UpdateUserSettings :exec
UPDATE
    "user"
SET
    region = $2,
    cf_limit = $3,
    dietary_requirement = $4
WHERE
    firebase_uid = $1;

-- name: GetUserByFirebaseUID :one
SELECT
    *
FROM
    "user"
WHERE
    firebase_uid = $1;

-- name: UpdateUserFCMToken :exec
UPDATE
    "user"
SET
    fcm_token = $2
WHERE
    firebase_uid = $1;

-- name: GetUserFCMTokens :many
SELECT
    fcm_token
FROM
    "user"
WHERE
    fcm_token IS NOT NULL;
