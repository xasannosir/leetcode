SELECT e.name AS name, b.bonus FROM Employee AS e LEFT JOIN Bonus AS b ON e.empId = b.empId WHERE e.empId NOT IN (SELECT empId FROM Bonus WHERE bonus >= 1000);
