package dao

import (
	"database/sql"
	"e-food/models"
	"fmt"
	"github.com/google/martian/log"
)

func GetUserCart(db *sql.DB, email string) (models.CartPreview, error) {
	q := fmt.Sprintf("SELECT p.productId,p.name, uc.totalQty,p.unitPrice, p.imageUrl FROM user_cart_item uc INNER JOIN product p where uc.productId = p.productId and uc.email='%s'", email)
	rows, err := db.Query(q)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	var cart models.CartPreview

	for rows.Next() {
		item := models.CartItem{}
		_ = rows.Scan(
			&item.ProductID,
			&item.ProductName,
			&item.Quantity,
			&item.UnitPrice,
			&item.ImageURL,
		)
		cart = append(cart, &item)
	}
	return cart, nil
}
