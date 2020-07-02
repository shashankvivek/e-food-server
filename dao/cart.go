package dao

import (
	"database/sql"
	"e-food/models"
	"fmt"
)

func GetGuestCart(db *sql.DB, sessionId string) (models.CartPreview, error) {
	q := fmt.Sprintf("SELECT p.productId,p.name, gc.totalQty,p.unitPrice, p.imageUrl FROM guest_cart_item gc INNER JOIN product p where gc.productId = p.productId and gc.sessionId='%s'", sessionId)
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	var cart models.CartPreview

	for rows.Next() {
		item := models.CartItem{}
		err = rows.Scan(
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

func AddItemToGuestCart(db *sql.DB, sessionId string, totalQty, productId int64) (bool, error) {
	res, err := db.Exec("INSERT INTO guest_cart_item (sessionId,totalQty,productId) VALUES (?, ?, ?)", sessionId, totalQty, productId)
	if err != nil {
		return false, err
	}
	insertedRow, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return insertedRow == 1, nil
}
