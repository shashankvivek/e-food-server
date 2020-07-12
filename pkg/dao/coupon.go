package dao

import (
	"crypto/rand"
	"database/sql"
	"encoding/base32"
	"errors"
	"fmt"
	"strings"
	"time"
)

func InsertNewCoupon(db *sql.DB, userLimit int, expTime time.Time, ruleSet string) (string, error) {
	randId := getToken(10)
	err := insertWithUniqueId(db, userLimit, expTime, ruleSet, randId)

	for err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		randId = getToken(10)
		err = insertWithUniqueId(db, userLimit, expTime, ruleSet, randId)
	}
	if err != nil {
		return "", err
	}
	return randId, nil
}

func CheckValidityOfCoupon(db *sql.DB, coupon string) error {
	row := db.QueryRow("SELECT userLimit,expiryDate from coupons where couponId = ? ", coupon)
	var userLimit int
	var expTime time.Time
	err := row.Scan(&userLimit, &expTime)
	if err != nil {
		return err
	}
	currentDate := time.Now().UTC()
	if expTime.After(currentDate) && userLimit > 0 {
		return nil
	} else {
		if userLimit < 1 {
			fmt.Println("User limit reached")
		}
		if expTime.Before(currentDate) {
			fmt.Println("coupon has expired")
		}
		return errors.New("invalid coupon")
	}
}

func insertWithUniqueId(db *sql.DB, userLimit int, expTime time.Time, ruleSet, randId string) error {
	_, err := db.Exec("INSERT INTO coupons (couponId, expiryDate, RuleSet, userLimit) VALUES (?,?,?,?)", randId, expTime, ruleSet, userLimit)
	if err != nil {
		return err
	}
	return nil
}

func getToken(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}
