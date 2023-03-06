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
