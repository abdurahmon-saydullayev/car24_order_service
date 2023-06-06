
CREATE OR REPLACE FUNCTION calculate_total_price()
RETURNS TRIGGER AS $$
BEGIN
    SELECT price_per_day INTO NEW.price_per_day
    FROM tarif
    WHERE id = NEW.tarif_id;
    
    SELECT discount_type INTO NEW.discount_type
    FROM discount
    WHERE id = NEW.discount_id;
    
    SELECT price_per_hour INTO NEW.mechanic_price_per_hour
    FROM mechanic
    WHERE id = NEW.mechanic_id;
    
    IF NEW.discount_type = 'fix' THEN
        NEW.total_price := (NEW.price_per_hour - NEW.discount_value) + (NEW.day_count * NEW.mechanic_price_per_hour);
    ELSIF NEW.discount_type = 'percentage' THEN
        NEW.total_price := (NEW.price_per_hour * (1 - NEW.discount_value / 100)) + (NEW.day_count * NEW.mechanic_price_per_hour);
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER calculate_total_price_trigger
BEFORE INSERT ON "order"
FOR EACH ROW
EXECUTE FUNCTION calculate_total_price();
