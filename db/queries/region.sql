-- name: CreateRegion :exec
INSERT INTO region (name)
    VALUES ($1);

-- name: ListRegions :many
SELECT
    name
FROM
    region;
