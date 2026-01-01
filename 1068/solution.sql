SELECT p.product_name AS product_name, s.year AS year, s.price AS price FROM Product AS p INNER JOIN Sales AS s ON p.product_id = s.product_id;
