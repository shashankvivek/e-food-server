package dao

import (
	"database/sql"
	"e-food/api/models"
	"e-food/pkg/integration"
	"errors"
	"fmt"
	"github.com/google/martian/log"
	"strings"
	"time"
)

func GetCustomerCart(db *sql.DB, email string) (models.CartPreview, string, error) {
	cartId, couponId, err := CreateOrGetCartDetails(db, email)
	if err != nil {
		fmt.Println(err.Error())
		return nil, "", err
	}
	rows, err := db.Query("SELECT p.productId,p.name,p.currency, ci.totalQty,p.unitPrice,\np.imageUrl FROM customer_cart_item ci ,  product p \nwhere ci.productId = p.productId and ci.cartId=?", cartId)
	if err != nil {
		log.Errorf(err.Error())
		return nil, "", err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		log.Errorf(err.Error())
		return nil, "", err
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
	return cart, couponId, nil
}

func AddItemToCustomerCart(db *sql.DB, email string, totalQty, productId int64) (*models.CartSuccessResponse, error) {
	cartId, _, err := CreateOrGetCartDetails(db, email)
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

	err = insertItemInUserCart(db, totalQty, cartId, productId)
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

func CreateOrGetCartDetails(db *sql.DB, email string) (int64, string, error) {
	_, err := db.Exec("INSERT into cart (email,couponId,createdAt) values (?,?,?)", email, "", time.Now().UTC())
	if err != nil && !strings.Contains(err.Error(), "Duplicate entry") {
		return -1, "", err
	}
	row := db.QueryRow("SELECT cartId, couponId from cart where email = ?", email)
	var cartId int64
	var couponId string
	err = row.Scan(&cartId, &couponId)
	if err != nil {
		return -1, "", err
	}
	return cartId, couponId, nil
}

func deleteExistingUserCartItemIfAny(db *sql.DB, cartId, productId int64) error {
	res, err := db.Exec("DELETE from customer_cart_item where productId = ? and cartId = ?", productId, cartId)
	if err != nil && !strings.Contains(err.Error(), "no row") {
		log.Errorf(err.Error())
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
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

func RemoveItemFromCustomerCart(db *sql.DB, productId int64, email string) error {
	cartId, _, err := CreateOrGetCartDetails(db, email)
	if err != nil {
		return err
	}
	err = deleteExistingUserCartItemIfAny(db, cartId, productId)
	if err != nil {
		return err
	}
	return nil
}

func ShiftGuestCartItemsToCustomer(db *sql.DB, guestHandler GuestCartHandler, sessionId, email string) error {
	guestCartItems, err := guestHandler.GetGuestCart(db, sessionId)
	if err != nil {
		return err
	}
	//look for each productId and email in user_cart_item table, if found "Update" else insert
	for _, gCartItem := range guestCartItems {
		//delete any prev entry of this product
		_ = RemoveItemFromCustomerCart(db, gCartItem.ProductID, email)
		_, err = AddItemToCustomerCart(db, email, gCartItem.Quantity, gCartItem.ProductID)
		if err != nil {
			return err
		}
	}
	//clean up guest_cart_item table
	_ = guestHandler.EmptyGuestCartItem(db, sessionId)
	return nil
}

func ApplyCouponToCart(db *sql.DB, coupon, email string) error {
	couponEntity, err := GetCouponDetails(db, coupon, email)
	if err != nil {
		return err
	}
	cartItems, _, err := GetCustomerCart(db, email)
	if err != nil {
		return err
	}
	productsFound := integration.CheckForMatchingProductsWithRuleSets(couponEntity.Rule.RuleSet, cartItems)
	if !productsFound {
		return errors.New("coupon condition not satisfied")
	}

	cartId, _, err := CreateOrGetCartDetails(db, email)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE cart SET couponId = ? where cartId = ? ", coupon, cartId)
	if err != nil {
		return err
	}
	return nil
}

func RemoveCouponFromCart(db *sql.DB, email string) error {
	_, err := db.Exec("UPDATE cart set couponId = '' where email= ?", email)
	if err != nil {
		return err
	}
	return nil
}

func RenewCart(db *sql.DB, cartId int64) error {
	_, err := db.Exec("DELETE  FROM customer_cart_item where cartId = ?", cartId)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE from cart where cartId = ?", cartId)
	if err != nil {
		return err
	}
	return nil
}

func GetAppliedCouponIdOnCart(db *sql.DB, cartId int64) (string, error) {
	row := db.QueryRow("SELECT couponId from cart where  cartId = ?", cartId)
	var couponId = ""
	err := row.Scan(&couponId)
	if err != nil {
		return "", err
	}
	return couponId, nil
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
