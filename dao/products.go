package dao

import (
	"database/sql"
	"e-food/models"
	"fmt"
)

func GetProductsBySubCategory(dbClient *sql.DB, scId int64) (models.Products, error) {
	q := fmt.Sprintf("SELECT productId,name,description,bcId,currency,imageUrl,discountPercentage,unitPrice,scId FROM ecommerce.product where scId=%d", scId)
	fmt.Println(q)
	rows, err := dbClient.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var retVal models.Products
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		product := &models.Product{}
		err = rows.Scan(
			&product.ProductID,
			&product.Name,
			&product.Description,
			&product.BcID,
			&product.Currency,
			&product.ImageURL,
			&product.DiscountPercentage,
			&product.UnitPrice,
			&product.ScID)
		if err != nil {
			return nil, err
		}
		retVal = append(retVal, product)
	}
	return retVal, nil
}
