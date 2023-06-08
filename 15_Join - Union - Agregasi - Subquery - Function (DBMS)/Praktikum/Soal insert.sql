INSERT INTO operators (operator_name) VALUES ('Operator A'), ('Operator B'), ('Operator C'), ('Operator D'), ('Operator E');
INSERT INTO product_types (type_name) VALUES ('Type A'), ('Type B'), ('Type C');
INSERT INTO products (product_name, product_type_id, operator_id) VALUES ('Product 1', 1, 3), ('Product 2', 1, 3);
INSERT INTO products (product_name, product_type_id, operator_id) VALUES ('Product 3', 2, 1), ('Product 4', 2, 1), ('Product 5', 2, 1);
INSERT INTO products (product_name, product_type_id, operator_id) VALUES ('Product 6', 3, 4), ('Product 7', 3, 4), ('Product 8', 3, 4);
UPDATE products SET product_description = 'Product 1 description' WHERE product_name = 'Product 1';
UPDATE products SET product_description = 'Product 2 description' WHERE product_name = 'Product 2';
UPDATE products SET product_description = 'Product 3 description' WHERE product_name = 'Product 3';
UPDATE products SET product_description = 'Product 4 description' WHERE product_name = 'Product 4';
UPDATE products SET product_description = 'Product 5 description' WHERE product_name = 'Product 5';
UPDATE products SET product_description = 'Product 6 description' WHERE product_name = 'Product 6';
UPDATE products SET product_description = 'Product 7 description' WHERE product_name = 'Product 7';
UPDATE products SET product_description = 'Product 8 description' WHERE product_name = 'Product 8';
INSERT INTO payment_methods (method_name) VALUES ('Method A'), ('Method B'), ('Method C');
INSERT INTO users (username, email) VALUES ('User A', 'usera@example.com'), ('User B', 'userb@example.com'), ('User C', 'userc@example.com'), ('User D', 'userd@example.com'), ('User E', 'usere@example.com');
-- User A
INSERT INTO transactions (user_id, payment_method_id) VALUES (1, 1), (1, 2), (1, 3);

-- User B
INSERT INTO transactions (user_id, payment_method_id) VALUES (2, 1), (2, 2), (2, 3);

-- User C
INSERT INTO transactions (user_id, payment_method_id) VALUES (3, 1), (3, 2), (3, 3);

-- User D
INSERT INTO transactions (user_id, payment_method_id) VALUES (4, 1), (4, 2), (4, 3);

-- User E
INSERT INTO transactions (user_id, payment_method_id) VALUES (5, 1), (5, 2), (5, 3);
-- User A, Transaction 1
INSERT INTO transaction_products (transaction_id, product_id) VALUES (1, 1), (1, 2), (1, 3);

-- User A, Transaction 2
INSERT INTO transaction_products (transaction_id, product_id) VALUES (2, 4), (2, 5), (2, 6);

-- User A, Transaction 3
INSERT INTO transaction_products (transaction_id, product_id) VALUES (3,
