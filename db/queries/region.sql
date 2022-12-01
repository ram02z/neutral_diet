-- name: CreateRegion :exec
INSERT INTO region (name)
    VALUES ($1);
