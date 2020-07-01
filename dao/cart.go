package dao

import (
	"database/sql"
	"fmt"
)

func GetGuestCartItemCount(db *sql.DB, sessionId string) (int64, error) {
	var count int64
	q := fmt.Sprintf("SELECT COUNT(*) AS Count FROM ecommerce.guest_cart_item where sessionId = '%s'", sessionId)
	err := db.QueryRow(q).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}
