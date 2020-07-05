package utils

import (
	"database/sql"
	"e-food/dao"
)

func ShiftGuestCartItemsToUserCart(db *sql.DB, sessionId, email string) error {
	guestCartItems, err := dao.GetGuestCart(db, sessionId)
	if err != nil {
		return err
	}
	//look for each productId and email in user_cart_item table, if found "Update" else insert
	for _, gCartItem := range guestCartItems {
		//delete any prev entry of this product
		_ = dao.RemoveItemFromUserCart(db, gCartItem.ProductID, email)
		_, err = dao.AddItemToUserCart(db, email, gCartItem.Quantity, gCartItem.ProductID)
		if err != nil {
			return err
		}
	}
	//clean up guest_cart_item table
	_ = dao.EmptyGuestCartItem(db, sessionId)
	return nil
}
