CREATE OR REPLACE PROCEDURE sp_create_product(
    name VARCHAR(50),
    description VARCHAR(255),
    stock INT,
    price DECIMAL(18, 2),
    category_id INT,
    user_id INT,
    pic_path TEXT[]
)
LANGUAGE plpgsql
AS $$
DECLARE
    product_id VARCHAR(255);
BEGIN
    SELECT fn_generate_product_id() INTO product_id;
    
    -- Insert product data
    INSERT INTO tb_products (
        name, 
        description, 
        stock, 
        price, 
        category_id, 
        user_id, 
        product_id
    )
    VALUES (
        name,
        description,
        stock,
        price,
        category_id,
        user_id,
        product_id
    );

    -- Insert picture paths as array
    INSERT INTO tb_pic_path(
        product_id,
        pic_path
    ) 
    VALUES (
        product_id,
        pic_path
    );

    -- Error handling
    EXCEPTION
        WHEN others THEN
            -- Re-raise the error
            RAISE EXCEPTION '%', SQLERRM;
END;
$$;

