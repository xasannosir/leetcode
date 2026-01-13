-- Write your PostgreSQL query statement below
SELECT user_id, name, mail FROM Users WHERE mail LIKE '%@leetcode.com' AND REGEXP_LIKE(LEFT(mail, LENGTH(mail) - 13), '^[a-zA-Z0-9._-]+$') AND REGEXP_LIKE(LEFT(mail, 1), '^[a-zA-Z]+$');

-- Write your MySQL query statement below
SELECT user_id, name, mail FROM Users WHERE BINARY RIGHT(mail, 13) = '@leetcode.com' AND REGEXP_LIKE(LEFT(mail, LENGTH(mail) - 13), '^[a-zA-Z0-9._-]+$') AND REGEXP_LIKE(LEFT(mail, 1), '^[a-zA-Z]+$');