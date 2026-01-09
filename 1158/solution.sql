-- Write your PostgreSQL query statement below
SELECT user_id AS buyer_id, join_date, (SELECT COUNT(*) FROM Orders WHERE buyer_id = user_id AND order_date::TEXT LIKE '2019%') AS orders_in_2019 FROM Users;

-- Write your MySQL query statement below
SELECT user_id AS buyer_id, join_date, (SELECT COUNT(*) AS counter FROM Orders WHERE buyer_id = user_id AND CAST(order_date AS CHAR) LIKE '2019%') AS orders_in_2019 FROM Users;