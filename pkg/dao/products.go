package dao

import (
	"database/sql"
	"e-food/api/models"
	"fmt"
	"github.com/google/martian/log"
)

type ProductHandler interface {
	GetUnitsInStock(db *sql.DB, productId int64) (int64, error)
	GetProductsBySubCategory(dbClient *sql.DB, scId int64) (models.Products, error)
}
type product struct {
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

func CreateProductHandler() ProductHandler {
	return &product{}
}

func (p *product) GetUnitsInStock(db *sql.DB, productId int64) (int64, error) {
	var unitsInStock = 0
	row := db.QueryRow("SELECT unitsInStock FROM ecommerce.product where productId=?", productId)

	err := row.Scan(&unitsInStock)
	if err != nil {
		return -1, err
	}
	return int64(unitsInStock), nil
}

func (p *product) GetProductsBySubCategory(dbClient *sql.DB, scId int64) (models.Products, error) {
	q := fmt.Sprintf("SELECT productId,name,description,bcId,currency,imageUrl,discountPercentage,unitPrice,scId,unitsInStock FROM ecommerce.product where scId=%d", scId)
	fmt.Println(q)
	rows, err := dbClient.Query(q)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	var retVal models.Products
	if err := rows.Err(); err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	for rows.Next() {
		product := product{}
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
			log.Errorf(err.Error())
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
