package dao

import (
	"database/sql"
	"e-food/constants"
	"e-food/models"
	"errors"
	"fmt"
	"strings"
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

func AddItemToGuestCart(db *sql.DB, sessionId string, totalQty, productId int64) (bool, string, error) {
	msg := ""
	unitsInStock, err := GetUnitsInStock(db, productId)
	if err != nil {
		return false, msg, err
	}
	if unitsInStock < 1 {
		return false, msg, errors.New("item is out of stock")
	}

	if totalQty > unitsInStock {
		totalQty = unitsInStock
		msg = "Reached max stock quantity"
	}
	// TODO: Update if already added
	itemQtyInCart, err := checkItemInGuestCart(db, sessionId, productId)
	if err != nil {
		return false, "", err
	}
	if itemQtyInCart > 0 {
		isSuccess, msg, err := updateGuestCartItemQuantity(db, unitsInStock, totalQty, itemQtyInCart, productId, sessionId)
		if err != nil {
			return false, msg, err
		}
		return isSuccess, msg, nil
	} else {
		isSuccess, err := insertGuestCartItem(db, unitsInStock, totalQty, productId, sessionId)
		if err != nil {
			return false, "", err
		}
		return isSuccess, msg, nil
	}
}

func RemoveItemFromGuestCart(db *sql.DB, productId int64, sessionId string) (bool, error) {
	res, err := db.Exec("DELETE from guest_cart_item where sessionId = ? and productId = ?", sessionId, productId)
	if err != nil {
		return false, err
	}
	defer db.Close()
	deletedRow, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return deletedRow == 1, nil
}

func checkItemInGuestCart(db *sql.DB, sessionId string, productId int64) (int64, error) {
	addedQty := 0
	row := db.QueryRow("SELECT totalQty from guest_cart_item where productId = ? and sessionId = ?", productId, sessionId)

	err := row.Scan(&addedQty)
	if err != nil && !strings.Contains(err.Error(), "no row") {
		return 0, err
	}
	return int64(addedQty), nil
}

func updateGuestCartItemQuantity(db *sql.DB, unitsInStock, totalQty, itemQtyInCart, productId int64, sessionId string) (bool, string, error) {
	msg := ""
	qtyToAdd := itemQtyInCart + totalQty
	if qtyToAdd > constants.MAX_ALLOWED_CART_ITEM_QTY {
		qtyToAdd = constants.MAX_ALLOWED_CART_ITEM_QTY
		msg = "Reached max stock quantity"
	}
	tx, err := db.Begin()
	if err != nil {
		return false, "", err
	}
	res, err := tx.Exec("UPDATE guest_cart_item set totalQty = ? where sessionId = ? and productId = ?", qtyToAdd, sessionId, productId)
	if err != nil {
		tx.Rollback()
		return false, "", err
	}
	updatedGuestCartItemRow, _ := res.RowsAffected()

	remainingUnitsInStock := unitsInStock - qtyToAdd
	res, err = tx.Exec("UPDATE product SET unitsInStock = ? WHERE (productId = ?)", remainingUnitsInStock, productId)
	if err != nil {
		tx.Rollback()
		return false, "", err
	}
	updatedProductRow, _ := res.RowsAffected()
	if updatedGuestCartItemRow == 1 && updatedProductRow == 1 {
		tx.Commit()
		return true, msg, nil
	} else {
		tx.Rollback()
		return false, "", errors.New("cart update transaction failed")
	}
}

func insertGuestCartItem(db *sql.DB, unitsInStock, totalQty, productId int64, sessionId string) (bool, error) {
	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	res, err := tx.Exec("INSERT INTO guest_cart_item (sessionId,totalQty,productId) VALUES (?, ?, ?)", sessionId, totalQty, productId)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	insertedRow, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	remainingUnitsInStock := unitsInStock - totalQty

	res, err = tx.Exec("UPDATE product SET unitsInStock = ? WHERE (productId = ?)", remainingUnitsInStock, productId)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	updatedRow, _ := res.RowsAffected()
	if insertedRow == 1 && updatedRow == 1 {
		tx.Commit()
		return true, nil
	} else {
		tx.Rollback()
		return false, errors.New("cart insert transaction failed")
	}
}
