SELECT transactions.*, products.product_name, users.username
FROM transactions
JOIN transaction_details ON transactions.id = transaction_details.transaction_id
JOIN products ON transaction_details.product_id = products.id
JOIN users ON transactions.user_id = users.id;
