-- Write your PostgreSQL query statement below
(SELECT s.season AS season, p.category AS category, SUM(s.total_quantity) AS total_quantity, SUM(s.total_revenue) AS total_revenue FROM products AS p INNER JOIN (SELECT 'Fall' AS season, s.product_id AS product_id, SUM(s.quantity) AS total_quantity, SUM(s.quantity::NUMERIC * s.price) AS total_revenue FROM sales AS s WHERE s.sale_date::TEXT LIKE '%-09-%' OR s.sale_date::TEXT LIKE '%-10-%' OR s.sale_date::TEXT LIKE '%-11-%' GROUP BY s.product_id) AS s ON s.product_id = p.product_id GROUP BY p.category, s.season ORDER BY total_quantity DESC, total_revenue DESC, category LIMIT 1)

UNION

(SELECT s.season AS season, p.category AS category, SUM(s.total_quantity) AS total_quantity, SUM(s.total_revenue) AS total_revenue FROM products AS p INNER JOIN (SELECT 'Spring' AS season, s.product_id AS product_id, SUM(s.quantity) AS total_quantity, SUM(s.quantity::NUMERIC * s.price) AS total_revenue FROM sales AS s WHERE s.sale_date::TEXT LIKE '%-03-%' OR s.sale_date::TEXT LIKE '%-04-%' OR s.sale_date::TEXT LIKE '%-05-%' GROUP BY s.product_id) AS s ON s.product_id = p.product_id GROUP BY p.category, s.season ORDER BY total_quantity DESC, total_revenue DESC, category LIMIT 1)

UNION

(SELECT s.season AS season, p.category AS category, SUM(s.total_quantity) AS total_quantity, SUM(s.total_revenue) AS total_revenue FROM products AS p INNER JOIN (SELECT 'Summer' AS season, s.product_id AS product_id, SUM(s.quantity) AS total_quantity, SUM(s.quantity::NUMERIC * s.price) AS total_revenue FROM sales AS s WHERE s.sale_date::TEXT LIKE '%-06-%' OR s.sale_date::TEXT LIKE '%-07-%' OR s.sale_date::TEXT LIKE '%-08-%' GROUP BY s.product_id) AS s ON s.product_id = p.product_id GROUP BY p.category, s.season ORDER BY total_quantity DESC, total_revenue DESC, category LIMIT 1)

UNION 

(SELECT s.season AS season, p.category AS category, SUM(s.total_quantity) AS total_quantity, SUM(s.total_revenue) AS total_revenue FROM products AS p INNER JOIN (SELECT 'Winter' AS season, s.product_id AS product_id, SUM(s.quantity) AS total_quantity, SUM(s.quantity::NUMERIC * s.price) AS total_revenue FROM sales AS s WHERE s.sale_date::TEXT LIKE '%-12-%' OR s.sale_date::TEXT LIKE '%-01-%' OR s.sale_date::TEXT LIKE '%-02-%' GROUP BY s.product_id) AS s ON s.product_id = p.product_id GROUP BY p.category, s.season ORDER BY total_quantity DESC, total_revenue DESC, category LIMIT 1)

ORDER BY season

-- Write your MySQL query statement below
(SELECT s.season AS season, p.category AS category, SUM(s.total_quantity) AS total_quantity, SUM(s.total_revenue) AS total_revenue FROM products AS p INNER JOIN (SELECT 'Fall' AS season, s.product_id AS product_id, SUM(s.quantity) AS total_quantity, SUM(s.quantity * s.price) AS total_revenue FROM sales AS s WHERE CAST(s.sale_date AS CHAR) LIKE '%-09-%' OR CAST(s.sale_date AS CHAR) LIKE '%-10-%' OR CAST(s.sale_date AS CHAR) LIKE '%-11-%' GROUP BY s.product_id) AS s ON s.product_id = p.product_id GROUP BY p.category, s.season ORDER BY total_quantity DESC, total_revenue DESC, category LIMIT 1)

UNION

(SELECT s.season AS season, p.category AS category, SUM(s.total_quantity) AS total_quantity, SUM(s.total_revenue) AS total_revenue FROM products AS p INNER JOIN (SELECT 'Spring' AS season, s.product_id AS product_id, SUM(s.quantity) AS total_quantity, SUM(s.quantity * s.price) AS total_revenue FROM sales AS s WHERE CAST(s.sale_date AS CHAR) LIKE '%-03-%' OR CAST(s.sale_date AS CHAR) LIKE '%-04-%' OR CAST(s.sale_date AS CHAR) LIKE '%-05-%' GROUP BY s.product_id) AS s ON s.product_id = p.product_id GROUP BY p.category, s.season ORDER BY total_quantity DESC, total_revenue DESC, category LIMIT 1)

UNION

(SELECT s.season AS season, p.category AS category, SUM(s.total_quantity) AS total_quantity, SUM(s.total_revenue) AS total_revenue FROM products AS p INNER JOIN (SELECT 'Summer' AS season, s.product_id AS product_id, SUM(s.quantity) AS total_quantity, SUM(s.quantity * s.price) AS total_revenue FROM sales AS s WHERE CAST(s.sale_date AS CHAR) LIKE '%-06-%' OR CAST(s.sale_date AS CHAR) LIKE '%-07-%' OR CAST(s.sale_date AS CHAR) LIKE '%-08-%' GROUP BY s.product_id) AS s ON s.product_id = p.product_id GROUP BY p.category, s.season ORDER BY total_quantity DESC, total_revenue DESC, category LIMIT 1)

UNION 

(SELECT s.season AS season, p.category AS category, SUM(s.total_quantity) AS total_quantity, SUM(s.total_revenue) AS total_revenue FROM products AS p INNER JOIN (SELECT 'Winter' AS season, s.product_id AS product_id, SUM(s.quantity) AS total_quantity, SUM(s.quantity * s.price) AS total_revenue FROM sales AS s WHERE CAST(s.sale_date AS CHAR) LIKE '%-12-%' OR CAST(s.sale_date AS CHAR) LIKE '%-01-%' OR CAST(s.sale_date AS CHAR) LIKE '%-02-%' GROUP BY s.product_id) AS s ON s.product_id = p.product_id GROUP BY p.category, s.season ORDER BY total_quantity DESC, total_revenue DESC, category LIMIT 1)

ORDER BY season