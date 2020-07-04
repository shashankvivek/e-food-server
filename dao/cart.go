package dao

import (
	"database/sql"
	"e-food/constants"
	"e-food/models"
	"errors"
	"fmt"
	"github.com/google/martian/log"
	"strings"
)

func GetGuestCart(db *sql.DB, sessionId string) (models.CartPreview, error) {
	q := fmt.Sprintf("SELECT p.productId,p.name, gc.totalQty,p.unitPrice, p.imageUrl FROM guest_cart_item gc INNER JOIN product p where gc.productId = p.productId and gc.sessionId='%s'", sessionId)
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

func AddItemToGuestCart(db *sql.DB, sessionId string, totalQty, productId int64) (*models.CartSuccessResponse, error) {
	msg := "Item added to cart"
	unitsInStock, err := GetUnitsInStock(db, productId)
	if err != nil {
		return nil, err
	}
	if unitsInStock < 1 {
		return nil, errors.New("item out of stock")
	}
	if totalQty > unitsInStock {
		totalQty = unitsInStock
		msg = "Reached max stock quantity"
	}
	itemQtyInCart, err := getItemQtyInGuestCart(db, sessionId, productId)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	if itemQtyInCart > 0 {
		updateMsg, addedQty, err := updateItemQtyInGuestCart(db, unitsInStock, totalQty, itemQtyInCart, productId, sessionId)
		if err != nil {
			log.Errorf(err.Error())
			return nil, err
		}
		var retVal = &models.CartSuccessResponse{
			Success:  true,
			Message:  updateMsg + msg,
			QtyAdded: addedQty,
		}
		return retVal, nil
	} else {
		err := insertItemInGuestCart(db, unitsInStock, totalQty, productId, sessionId)
		if err != nil {
			return nil, err
		}
		var retVal = &models.CartSuccessResponse{
			Success:  true,
			Message:  msg,
			QtyAdded: totalQty,
		}
		return retVal, nil
	}
}

func RemoveItemFromGuestCart(db *sql.DB, productId int64, sessionId string) (bool, error) {
	itemQtyInCart, err := getItemQtyInGuestCart(db, sessionId, productId)
	if err != nil {
		log.Errorf(err.Error())
		return false, err
	}
	if itemQtyInCart < 1 {
		return false, errors.New("item does not exist")
	}
	exitingUnitsInStock, err := GetUnitsInStock(db, productId)
	if err != nil {
		log.Errorf(err.Error())
		return false, err
	}
	finalQty := exitingUnitsInStock + itemQtyInCart
	tx, err := db.Begin()
	if err != nil {
		return false, err
	}

	res, err := tx.Exec("DELETE from guest_cart_item where sessionId = ? and productId = ?", sessionId, productId)
	if err != nil {
		return false, err
	}

	deletedRow, _ := res.RowsAffected()
	res, err = tx.Exec("UPDATE product SET unitsInStock = ? WHERE (productId = ?)", finalQty, productId)
	if err != nil {
		tx.Rollback()
		return false, err
	}
	updatedRow, _ := res.RowsAffected()
	if deletedRow == 1 && updatedRow == 1 {
		tx.Commit()
		return true, nil
	} else {
		tx.Rollback()
		return false, errors.New("cart item removal transaction error")
	}
}

func getItemQtyInGuestCart(db *sql.DB, sessionId string, productId int64) (int64, error) {
	addedQty := 0
	row := db.QueryRow("SELECT totalQty from guest_cart_item where productId = ? and sessionId = ?", productId, sessionId)
	err := row.Scan(&addedQty)
	if err != nil && !strings.Contains(err.Error(), "no row") {
		log.Errorf(err.Error())
		return 0, err
	}
	return int64(addedQty), nil
}

func updateItemQtyInGuestCart(
	db *sql.DB,
	unitsInStock, totalQty, itemQtyInCart, productId int64,
	sessionId string) (string, int64, error) {

	msg := "Item added successfully. "
	qtyToAdd := itemQtyInCart + totalQty
	if qtyToAdd > constants.MAX_ALLOWED_CART_ITEM_QTY {
		qtyToAdd = constants.MAX_ALLOWED_CART_ITEM_QTY
		msg = "Item reached max stock quantity. "
	}
	tx, err := db.Begin()
	if err != nil {
		log.Errorf(err.Error())
		return "", -1, err
	}
	res, err := tx.Exec("UPDATE guest_cart_item set totalQty = ? where sessionId = ? and productId = ?", qtyToAdd, sessionId, productId)
	if err != nil {
		tx.Rollback()
		log.Errorf(err.Error())
		return "", -1, err
	}
	updatedGuestCartItemRow, _ := res.RowsAffected()

	remainingUnitsInStock := unitsInStock - qtyToAdd
	res, err = tx.Exec("UPDATE product SET unitsInStock = ? WHERE (productId = ?)", remainingUnitsInStock, productId)
	if err != nil {
		tx.Rollback()
		log.Errorf(err.Error())
		return "", -1, err
	}
	updatedProductRow, _ := res.RowsAffected()
	if updatedGuestCartItemRow == 1 && updatedProductRow == 1 {
		tx.Commit()
		return msg, qtyToAdd, nil
	} else {
		tx.Rollback()
		return "", -1, errors.New("cart update transaction failed")
	}
}

func insertItemInGuestCart(db *sql.DB, unitsInStock, totalQty, productId int64, sessionId string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec("INSERT INTO guest_cart_item (sessionId,totalQty,productId) VALUES (?, ?, ?)", sessionId, totalQty, productId)
	if err != nil {
		tx.Rollback()
		return err
	}
	insertedRow, err := res.RowsAffected()
	if err != nil {
		return err
	}
	remainingUnitsInStock := unitsInStock - totalQty

	res, err = tx.Exec("UPDATE product SET unitsInStock = ? WHERE (productId = ?)", remainingUnitsInStock, productId)
	if err != nil {
		tx.Rollback()
		log.Errorf(err.Error())
		return err
	}
	updatedRow, _ := res.RowsAffected()
	if insertedRow == 1 && updatedRow == 1 {
		tx.Commit()
		return nil
	} else {
		tx.Rollback()
		return errors.New("cart insert transaction failed")
	}
}
