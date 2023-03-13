-- name: AddCarbonFootprintGoal :one
INSERT INTO "carbon_footprint_goal" (user_id, description, start_date, end_date, start_carbon_footprint, target_carbon_footprint)
    VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    id;

-- name: UpdateCarbonFootprintGoal :exec
UPDATE
    "carbon_footprint_goal"
SET
    completed = $3
WHERE
    user_id = $1
    AND id = $2;

-- name: DeleteCarbonFootprintGoal :exec
DELETE FROM "carbon_footprint_goal"
WHERE user_id = $1
    AND id = $2;

-- name: GetCarbonFootprintGoals :many
SELECT
    *
FROM
    "carbon_footprint_goal"
WHERE
    user_id = $1;
