CREATE OR REPLACE FUNCTION calculate_order_total_price()
    RETURNS TRIGGER AS $$
    DECLARE
        discount_value DECIMAL;
        mechanic_price DECIMAL;
        order_price DECIMAL;
    BEGIN
        SELECT price_per_day INTO order_price
        FROM tarif
        WHERE id = NEW.tarif_id;
        
        SELECT price_per_hour INTO mechanic_price
        FROM mechanic
        WHERE id = NEW.mechanic_id;
        
        order_price := (order_price * NEW.day_count) + mechanic_price - NEW.paid_price;
        
        IF NEW.discount IS NOT NULL THEN
            SELECT CASE
                WHEN discount_type = 'fix' THEN discount_amount
                WHEN discount_type = 'percentage' THEN order_price * (discount_amount / 100)
                ELSE 0
            END INTO discount_value
            FROM discount
            WHERE id = NEW.discount;
            
            order_price := order_price - discount_value;
        END IF;
        
        NEW.total_price := order_price;
        
        UPDATE mechanic
        SET status = TRUE
        WHERE id = NEW.mechanic_id;
        
        UPDATE car
        SET status = TRUE
        WHERE id = NEW.car_id;
        
        RETURN NEW;
    END;
    $$ LANGUAGE plpgsql;
