CREATE OR REPLACE PROCEDURE sp_create_product(
   name VARCHAR(50),
   description VARCHAR(255),
   stock INT,
   price DECIMAL(18, 2),
   category_id INT,
   user_id INT,
   INOUT v_pic_path TEXT[],
   INOUT out_product_id VARCHAR(255)
)
LANGUAGE plpgsql
AS $$
BEGIN
   -- Generate product ID
   SELECT fn_generate_product_id() INTO out_product_id;

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
       out_product_id
   );

   -- Insert picture paths
   INSERT INTO tb_pic_path (
       product_id,
       pic_path
   ) 
   VALUES (
       out_product_id,
       v_pic_path 
   );

   -- Return inserted paths
   SELECT pic_path
   INTO v_pic_path
   FROM tb_pic_path
   WHERE product_id = out_product_id;

EXCEPTION
   WHEN OTHERS THEN
       -- Rollback transaction
       v_pic_path := NULL;
       out_product_id := NULL;
       RAISE EXCEPTION 'Error occurred while creating product: %', SQLERRM;
END;
$$;

