-- Write your PostgreSQL query statement below
SELECT ROUND(SUM(tiv_2016)::NUMERIC, 2) AS tiv_2016 FROM Insurance WHERE pid IN (SELECT a1.pid AS pid FROM Insurance AS a1 INNER JOIN Insurance AS a2 ON a1.tiv_2015 = a2.tiv_2015 AND a1.pid <> a2.pid) AND pid NOT IN (SELECT a1.pid AS pid FROM Insurance AS a1 INNER JOIN Insurance AS a2 ON a1.lat = a2.lat AND a1.lon = a2.lon AND a1.pid <> a2.pid);

-- Write your MySQL query statement below
SELECT ROUND(SUM(tiv_2016), 2) AS tiv_2016 FROM Insurance WHERE pid IN (SELECT a1.pid AS pid FROM Insurance AS a1 INNER JOIN Insurance AS a2 ON a1.tiv_2015 = a2.tiv_2015 AND a1.pid <> a2.pid) AND pid NOT IN (SELECT a1.pid AS pid FROM Insurance AS a1 INNER JOIN Insurance AS a2 ON a1.lat = a2.lat AND a1.lon = a2.lon AND a1.pid <> a2.pid);