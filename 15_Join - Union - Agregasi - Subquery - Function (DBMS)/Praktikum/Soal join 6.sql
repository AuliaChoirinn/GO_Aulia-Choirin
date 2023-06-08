CREATE FUNCTION delete_transaction(transaction_id INT)
RETURNS VOID
AS $$
BEGIN
  DELETE FROM transaction_details WHERE transaction_id = $1;
  DELETE FROM transactions WHERE id = $1;
END;
$$ LANGUAGE plpgsql;
