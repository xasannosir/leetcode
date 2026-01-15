-- Write your PostgreSQL query statement below
SELECT r.contest_id AS contest_id, ROUND((COUNT(r.user_id)::NUMERIC/(SELECT COUNT(*) FROM Users)::NUMERIC)*100, 2) AS percentage FROM Users AS u INNER JOIN Register AS r ON u.user_id = r.user_id GROUP BY r.contest_id ORDER BY percentage DESC, contest_id ASC;

-- Write your MySQL query statement below
SELECT r.contest_id AS contest_id, ROUND((COUNT(r.user_id)/(SELECT COUNT(*) FROM Users))*100, 2) AS percentage FROM Users AS u INNER JOIN Register AS r ON u.user_id = r.user_id GROUP BY r.contest_id ORDER BY percentage DESC, contest_id ASC;