-- Write your PostgreSQL query statement below
SELECT s.score AS score, o.rank AS rank FROM Scores AS s LEFT JOIN (SELECT score, ROW_NUMBER() OVER(ORDER BY score DESC) AS rank FROM Scores GROUP BY score) AS o ON o.score = s.score ORDER BY s.score DESC;

-- Write your MySQL query statement below
SELECT s.score AS score, o.rank AS 'rank' FROM Scores AS s LEFT JOIN (SELECT score, ROW_NUMBER() OVER(ORDER BY score DESC) AS 'rank' FROM Scores GROUP BY score) AS o ON o.score = s.score ORDER BY s.score DESC;