package dao

import (
	"database/sql"
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
	err = insertItemInGuestCart(db, unitsInStock, totalQty, productId, sessionId)
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

func RemoveItemFromGuestCart(db *sql.DB, productId int64, sessionId string) error {
	itemQtyInCart, err := GetItemQtyInGuestCart(db, sessionId, productId)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}
	if itemQtyInCart < 1 {
		return errors.New("item does not exist")
	}
	res, err := db.Exec("DELETE from guest_cart_item where sessionId = ? and productId = ?", sessionId, productId)
	if err != nil {
		return err
	}
	deletedRow, _ := res.RowsAffected()
	if deletedRow == 1 {
		return nil
	} else {
		return errors.New("error removing item from guest cart")
	}
}

func GetItemQtyInGuestCart(db *sql.DB, sessionId string, productId int64) (int64, error) {
	addedQty := 0
	row := db.QueryRow("SELECT totalQty from guest_cart_item where productId = ? and sessionId = ?", productId, sessionId)
	err := row.Scan(&addedQty)
	if err != nil && !strings.Contains(err.Error(), "no row") {
		log.Errorf(err.Error())
		return 0, err
	}
	return int64(addedQty), nil
}

func insertItemInGuestCart(db *sql.DB, unitsInStock, totalQty, productId int64, sessionId string) error {
	res, err := db.Exec("INSERT INTO guest_cart_item (sessionId,totalQty,productId) VALUES (?, ?, ?)", sessionId, totalQty, productId)
	if err != nil {
		return err
	}
	insertedRow, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if insertedRow == 1 {
		return nil
	} else {
		return errors.New("cart insert transaction failed")
	}
}

func EmptyGuestCartItem(db *sql.DB, sessionId string) error {
	row, err := db.Exec("DELETE from guest_cart_item where sessionId = ? ", sessionId)
	if err != nil {
		return err
	}
	fmt.Print("Delete Guest Cart Item Count: ")
	fmt.Println(row.RowsAffected())
	fmt.Println("=============================")
	return nil
}
