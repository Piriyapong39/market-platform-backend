package product

import (
	"fmt"
	"time"

	db "github.com/piriyapong39/market-platform/config"

	// import service
	"github.com/lib/pq"
)

type Product struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
	CategoryID  int      `json:"category_id"`
	PicPath     []string `json:"pic_path"`
	User_id     int      `json:"user_id"`
}

func _createProduct(product Product) ([]string, error) {
	db, err := db.Connection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// pic path format
	var configPicPath []string
	var productId string
	if err := db.QueryRow(`SELECT fn_generate_product_id() As product_id`).Scan(&productId); err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	now := time.Now()
	formatNow := now.Format("2006-01-02")
	for _, picName := range product.PicPath {
		realPath := fmt.Sprintf("./upload/product-pictures/%s/%v/%s/%s", formatNow, product.CategoryID, productId, picName)
		configPicPath = append(configPicPath, realPath)
	}
	picPaths := pq.Array(configPicPath)

	query := `CALL sp_create_product($1, $2, $3, $4, $5, $6, $7, $8);`
	_, err = db.Exec(
		query,
		product.Name,
		product.Description,
		product.Stock,
		product.Price,
		product.CategoryID,
		product.User_id,
		productId,
		picPaths,
	)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return configPicPath, nil
}
