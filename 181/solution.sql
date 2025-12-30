SELECT e.name AS Employee FROM Employee AS e INNER JOIN Employee AS m ON e.managerId = m.id WHERE e.salary > m.salary;
