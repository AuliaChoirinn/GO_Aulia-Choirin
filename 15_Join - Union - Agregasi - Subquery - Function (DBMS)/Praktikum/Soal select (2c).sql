SELECT *
FROM users
WHERE gender = 'F' AND created_at >= DATE(NOW()) - INTERVAL 7 DAY AND username LIKE '%a%'
