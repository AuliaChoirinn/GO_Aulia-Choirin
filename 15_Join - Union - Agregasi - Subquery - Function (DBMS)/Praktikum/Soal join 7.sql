CREATE FUNCTION update_total_qty(transaction_id INT)
RETURNS VOID
AS $$
BEGIN
  UPDATE transactions
  SET total_qty = (
    SELECT SUM(qty)
    FROM transaction_details
    WHERE transaction_id = $1
  )
  WHERE id = $1;
END;
$$ LANGUAGE plpgsql;
