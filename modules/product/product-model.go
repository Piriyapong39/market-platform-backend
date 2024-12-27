package product

import (
	"fmt"

	db "github.com/piriyapong39/market-platform/config"

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

func _createProduct(product Product) (string, error) {
	db, err := db.Connection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var outProductID string
	var outProductPicPath string
	picPaths := pq.Array(product.PicPath)

	query := `CALL sp_create_product($1, $2, $3, $4, $5, $6, $7, $8);`
	row := db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Stock,
		product.Price,
		product.CategoryID,
		product.User_id,
		picPaths,
		nil,
	)

	err = row.Scan(&outProductPicPath, &outProductID)
	if err != nil {
		return "", fmt.Errorf("failed to get product ID: %v", err)
	}

	// var picPathsArr []string

	return outProductID, nil
}
