-- Write your PostgreSQL query statement below
SELECT s.user_id AS user_id, ROUND((CASE WHEN COUNT(*) = 0 THEN 0.00 ELSE SUM(CASE WHEN action = 'confirmed' THEN 1 ELSE 0 END)::NUMERIC/COUNT(*) END)::NUMERIC, 2) AS confirmation_rate FROM Signups AS s LEFT JOIN Confirmations AS c ON s.user_id = c.user_id GROUP BY s.user_id;

-- Write your MySQL query statement below
SELECT s.user_id AS user_id, ROUND((CASE WHEN COUNT(*) = 0 THEN 0.00 ELSE SUM(CASE WHEN action = 'confirmed' THEN 1 ELSE 0 END)/COUNT(*) END), 2) AS confirmation_rate FROM Signups AS s LEFT JOIN Confirmations AS c ON s.user_id = c.user_id GROUP BY s.user_id;