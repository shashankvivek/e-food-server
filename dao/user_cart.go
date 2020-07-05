package dao

import (
	"database/sql"
	"e-food/models"
	"errors"
	"fmt"
	"github.com/google/martian/log"
	"strings"
)

func GetUserCart(db *sql.DB, email string) (models.CartPreview, error) {
	q := fmt.Sprintf("SELECT p.productId,p.name,p.currency, uc.totalQty,p.unitPrice, p.imageUrl FROM user_cart_item uc INNER JOIN product p where uc.productId = p.productId and uc.email='%s'", email)
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
			&item.Currency,
			&item.Quantity,
			&item.UnitPrice,
			&item.ImageURL,
		)
		cart = append(cart, &item)
	}
	return cart, nil
}

func AddItemToUserCart(db *sql.DB, email string, totalQty, productId int64) (*models.CartSuccessResponse, error) {
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
	err = insertItemInUserCart(db, totalQty, productId, email)
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

func insertItemInUserCart(db *sql.DB, totalQty, productId int64, email string) error {
	res, err := db.Exec("INSERT INTO user_cart_item (email,totalQty,productId) VALUES (?, ?, ?)", email, totalQty, productId)
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
		return errors.New("adding item to user cart failed")
	}
}

func RemoveItemFromUserCart(db *sql.DB, productId int64, email string) error {
	itemQtyInCart, err := GetItemQtyInUserCart(db, email, productId)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}
	if itemQtyInCart < 1 {
		return errors.New("item does not exist")
	}
	res, err := db.Exec("DELETE from user_cart_item where email = ? and productId = ?", email, productId)
	if err != nil {
		return err
	}
	deletedRow, _ := res.RowsAffected()
	if deletedRow == 1 {
		return nil
	} else {
		return errors.New("error removing item from User cart")
	}
}

func GetItemQtyInUserCart(db *sql.DB, email string, productId int64) (int64, error) {
	addedQty := 0
	row := db.QueryRow("SELECT totalQty from user_cart_item where productId = ? and email = ?", productId, email)
	err := row.Scan(&addedQty)
	if err != nil && !strings.Contains(err.Error(), "no row") {
		log.Errorf(err.Error())
		return 0, err
	}
	return int64(addedQty), nil
}

//TODO: use this logic when the order is being created
//func insertItemInUserCart(db *sql.DB, unitsInStock, totalQty, productId int64, email string) error {
//	tx, err := db.Begin()
//	if err != nil {
//		return err
//	}
//	res, err := tx.Exec("INSERT INTO user_cart_item (email,totalQty,productId) VALUES (?, ?, ?)", email, totalQty, productId)
//	if err != nil {
//		tx.Rollback()
//		return err
//	}
//	insertedRow, err := res.RowsAffected()
//	if err != nil {
//		return err
//	}
//	remainingUnitsInStock := unitsInStock - totalQty
//
//	res, err = tx.Exec("UPDATE product SET unitsInStock = ? WHERE (productId = ?)", remainingUnitsInStock, productId)
//	if err != nil {
//		tx.Rollback()
//		log.Errorf(err.Error())
//		return err
//	}
//	updatedRow, _ := res.RowsAffected()
//	if insertedRow == 1 && updatedRow == 1 {
//		tx.Commit()
//		return nil
//	} else {
//		tx.Rollback()
//		return errors.New("cart insert transaction failed")
//	}
//}
