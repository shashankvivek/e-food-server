package dao

import (
	"database/sql"
	"e-food/models"
	"fmt"
)

type Product struct {
	BcID               int64
	Currency           string
	Description        string
	DiscountPercentage float64
	ImageURL           string
	IsAvailable        bool
	Name               string
	ProductID          int64
	ScID               int64
	Sku                string
	UnitPrice          float64
	UnitsInStock       int
}

func GetProductsBySubCategory(dbClient *sql.DB, scId int64) (models.Products, error) {
	q := fmt.Sprintf("SELECT productId,name,description,bcId,currency,imageUrl,discountPercentage,unitPrice,scId,unitsInStock FROM ecommerce.product where scId=%d", scId)
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
		product := Product{}
		err = rows.Scan(
			&product.ProductID,
			&product.Name,
			&product.Description,
			&product.BcID,
			&product.Currency,
			&product.ImageURL,
			&product.DiscountPercentage,
			&product.UnitPrice,
			&product.ScID,
			&product.UnitsInStock)
		if err != nil {
			return nil, err
		}
		product.IsAvailable = product.UnitsInStock > 0 // or create a DB trigger to manage this flag in DB itself
		retVal = append(retVal, &models.Product{
			BcID:               product.BcID,
			Currency:           product.Currency,
			Description:        product.Description,
			DiscountPercentage: product.DiscountPercentage,
			ImageURL:           product.ImageURL,
			IsAvailable:        product.IsAvailable,
			Name:               product.Name,
			ProductID:          product.ProductID,
			ScID:               product.ScID,
			Sku:                product.Sku,
			UnitPrice:          product.UnitPrice,
		})
	}
	return retVal, nil
}
