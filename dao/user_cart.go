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
	_, err := createOrGetCartId(db, email)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	rows, err := db.Query("SELECT p.productId,p.name,p.currency, ci.totalQty,p.unitPrice, p.imageUrl FROM customer_cart_item ci ,  product p , cart c where ci.productId = p.productId and c.email=?", email)
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
	cartId, err := createOrGetCartId(db, email)
	if err != nil {
		return nil, err
	}
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

	err = deleteExistingUserCartItemIfAny(db, cartId, productId)
	if err != nil {
		return nil, err
	}

	err = insertItemInUserCart(db, totalQty, productId, cartId)
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

func createOrGetCartId(db *sql.DB, email string) (int64, error) {
	_, err := db.Exec("INSERT into cart (email) values (?)", email)
	if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
		return -1, err
	}
	row := db.QueryRow("SELECT cartId from cart where email = ?", email)
	var cartId int64
	err = row.Scan(&cartId)
	if err != nil {
		return -1, err
	}
	return cartId, nil
}

func deleteExistingUserCartItemIfAny(db *sql.DB, cartId, productId int64) error {
	res, err := db.Exec("DELETE from customer_cart_item where productId = ? and cartId = ?", productId, cartId)
	if err != nil && !strings.Contains(err.Error(), "no row") {
		log.Errorf(err.Error())
		return err
	}
	deletedRow, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if deletedRow == 1 || deletedRow == 0 {
		return nil
	} else {
		return errors.New("found more than 1 item to delete")
	}
}

func insertItemInUserCart(db *sql.DB, totalQty, cartId, productId int64) error {
	res, err := db.Exec("INSERT INTO customer_cart_item (cartId,totalQty,productId) VALUES (?, ?, ?)", cartId, totalQty, productId)
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
	cartId, err := createOrGetCartId(db, email)
	if err != nil {
		return err
	}
	res, err := db.Exec("DELETE from customer_cart_item where cartId = ? and productId = ?", cartId, productId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
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
