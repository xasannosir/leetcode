-- Write your PostgreSQL query statement below
SELECT a.id AS id FROM Weather AS a INNER JOIN Weather AS b ON (a.recordDate - b.recordDate) = 1 AND a.temperature > b.temperature;

-- Write your MySQL query statement below
SELECT a.id AS id FROM Weather AS a INNER JOIN Weather AS b ON DATEDIFF(a.recordDate, b.recordDate) = 1 AND a.temperature > b.temperature;