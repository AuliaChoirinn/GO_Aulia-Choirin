SELECT COUNT(*) AS total_transactions
FROM transactions
JOIN transaction_details ON transactions.id = transaction_details.transaction_id
JOIN products ON transaction_details.product_id = products.id
WHERE products.product_type_id = 2;
