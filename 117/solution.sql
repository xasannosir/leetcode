CREATE OR REPLACE FUNCTION NthHighestSalary(N INT) RETURNS TABLE (Salary INT) AS $$
BEGIN
    RETURN QUERY (
        -- Write your PostgreSQL query statement below.
        SELECT e.salary FROM (SELECT DISTINCT e.salary FROM Employee AS e ORDER BY e.salary DESC) AS e LIMIT (CASE WHEN N > 0 THEN 1 ELSE 0 END) OFFSET (CASE WHEN N > 0 THEN N-1 ELSE 0 END)
    );
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    DECLARE N1 INT;
    SET N1 = (N -1);

    RETURN (
        -- Write your MySQL query statement below.
        SELECT e.salary AS salary FROM (SELECT DISTINCT e.salary FROM Employee AS e ORDER BY e.salary DESC) AS e LIMIT 1 OFFSET N1

    );
END