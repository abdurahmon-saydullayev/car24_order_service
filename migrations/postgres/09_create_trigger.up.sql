CREATE TRIGGER calculate_order_total_price_trigger
    BEFORE INSERT ON "order"
    FOR EACH ROW
    EXECUTE FUNCTION calculate_order_total_price();
