-- Write your PostgreSQL query statement below
(SELECT u.name AS results FROM MovieRating AS mr INNER JOIN Users AS u ON u.user_id = mr.user_id GROUP BY mr.user_id, u.name HAVING COUNT(*) = (SELECT COUNT(*) AS max_count FROM MovieRating GROUP BY user_id ORDER BY max_count DESC LIMIT 1) ORDER BY u.name LIMIT 1)
UNION ALL
(SELECT m.title AS results FROM MovieRating AS mr INNER JOIN Movies AS m ON m.movie_id = mr.movie_id WHERE mr.created_at::TEXT LIKE '2020-02%' GROUP BY mr.movie_id, m.title HAVING SUM(mr.rating)::NUMERIC/COUNT(*)::NUMERIC = (SELECT SUM(rating)::NUMERIC/COUNT(*)::NUMERIC AS max_rating FROM MovieRating WHERE created_at::TEXT LIKE '2020-02%' GROUP BY movie_id ORDER BY max_rating DESC LIMIT 1) ORDER BY m.title LIMIT 1)

-- Write your MySQL query statement below
(SELECT u.name AS results FROM MovieRating AS mr INNER JOIN Users AS u ON u.user_id = mr.user_id GROUP BY mr.user_id, u.name HAVING COUNT(*) = (SELECT COUNT(*) AS max_count FROM MovieRating GROUP BY user_id ORDER BY max_count DESC LIMIT 1) ORDER BY u.name LIMIT 1)
UNION ALL
(SELECT m.title AS results FROM MovieRating AS mr INNER JOIN Movies AS m ON m.movie_id = mr.movie_id WHERE CAST(mr.created_at AS CHAR) LIKE '2020-02%' GROUP BY mr.movie_id, m.title HAVING SUM(mr.rating)/COUNT(*) = (SELECT SUM(rating)/COUNT(*) AS max_rating FROM MovieRating WHERE CAST(created_at AS CHAR) LIKE '2020-02%' GROUP BY movie_id ORDER BY max_rating DESC LIMIT 1) ORDER BY m.title LIMIT 1)