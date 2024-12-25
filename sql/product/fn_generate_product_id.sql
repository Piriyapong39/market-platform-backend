CREATE OR REPLACE FUNCTION fn_generate_product_id()
RETURNS TEXT AS $$
DECLARE
    product_id TEXT;
    curr_date TEXT;
    last_id INT;
BEGIN
    -- create current date
    curr_date := TO_CHAR(CURRENT_DATE, 'YYMMDD');
    
    -- select last id from tb_products
    -- Adding COALESCE to handle case when table is empty
    SELECT COALESCE(MAX(id), 0) INTO last_id
    FROM tb_products;
    
    -- combine 'P', date and last id + 1 together
    product_id := 'P' || curr_date || (last_id + 1);
    
    RETURN product_id;
END;
$$ LANGUAGE plpgsql;
