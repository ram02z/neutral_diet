-- name: AddDevice :exec
INSERT INTO "device" (user_id, fcm_token)
    VALUES ($1, $2)
ON CONFLICT (fcm_token)
    DO UPDATE SET
        user_id = $1;

-- name: GetDevices :many
SELECT
    *
FROM
    "device";

-- name: DeleteDeviceByUser :exec
DELETE FROM "device"
WHERE user_id = $1;

-- name: DeleteDevice :batchexec
DELETE FROM "device"
WHERE id = $1;
