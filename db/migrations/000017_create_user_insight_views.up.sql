CREATE VIEW daily_user_average_by_dietary_requirement AS
SELECT
    (SUM(f.carbon_footprint) / count(DISTINCT f.log_date))::DECIMAL AS average_carbon_footprint,
    u.dietary_requirement
FROM
    food_item_log f
    INNER JOIN "user" u ON f.user_id = u.id
GROUP BY
    u.dietary_requirement;

CREATE VIEW daily_user_average AS
SELECT
    (SUM(f.carbon_footprint) / count(DISTINCT f.log_date))::DECIMAL AS average_carbon_footprint
FROM
    food_item_log f;

CREATE VIEW food_item_log_user_streak AS
WITH dates (
    date
) AS (
    SELECT DISTINCT
        log_date,
        user_id
    FROM
        food_item_log
),
GROUPS AS (
    SELECT
        user_id,
        ROW_NUMBER() OVER (ORDER BY date) AS rn,
    date + - ROW_NUMBER() OVER (ORDER BY date) * INTERVAL '1 day' AS grp,
    date
FROM
    dates
)
SELECT
    user_id,
    COUNT(*) AS consecutive_dates,
    MIN(date) AS start_date,
    MAX(date) AS end_date,
    CURRENT_DATE < MAX(date) + 1 AS active
FROM
    GROUPS
GROUP BY
    user_id,
    grp
HAVING
    COUNT(*) > 1;
