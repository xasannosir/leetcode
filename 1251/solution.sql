-- Write your MySQL query statement below
SELECT p.product_id AS product_id, COALESCE(ROUND(SUM(p.price * us.units) / SUM(us.units), 2), 0.00) AS average_price FROM Prices AS p LEFT JOIN UnitsSold AS us ON p.product_id = us.product_id AND us.purchase_date >= p.start_date AND us.purchase_date <= p.end_date GROUP B p.product_id;

-- Write your PostgreSQL query statement below
SELECT p.product_id AS product_id, COALESCE(ROUND(SUM(p.price * us.units)::NUMERIC / SUM(us.units)::NUMERIC, 2), 0.00) AS average_price FROM Prices AS p LEFT JOIN UnitsSold AS us ON p.product_id = us.product_id AND us.purchase_date >= p.start_date AND us.purchase_date <= p.end_date GROUP BYp.product_id;