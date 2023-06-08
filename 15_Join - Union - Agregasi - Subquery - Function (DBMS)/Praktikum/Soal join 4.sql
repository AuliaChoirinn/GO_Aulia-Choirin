SELECT products.*, product_types.product_type_name
FROM products
JOIN product_types ON products.product_type_id = product_types.id;
