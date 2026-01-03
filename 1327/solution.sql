-- Write your PostgreSQL query statement below
SELECT p.product_name AS product_name, o.sum_unit AS unit FROM Products AS p INNER JOIN (SELECT product_id, SUM(unit) AS sum_unit FROM Orders WHERE CAST(order_date AS TEXT) LIKE '2020-02%' GROUP BY product_id HAVING SUM(unit) >= 100) AS o ON p.product_id = o.product_id;

-- Write your MySQL query statement below
SELECT p.product_name AS product_name, o.sum_unit AS unit FROM Products AS p INNER JOIN (SELECT product_id, SUM(unit) AS sum_unit FROM Orders WHERE CAST(order_date AS CHAR) LIKE '2020-02%' GROUP BY product_id HAVING SUM(unit) >= 100) AS o ON p.product_id = o.product_id;
