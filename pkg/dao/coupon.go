package dao

import (
	"crypto/rand"
	"database/sql"
	"e-food/pkg/entities"
	"encoding/base32"
	"encoding/json"
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

func ReduceUserLimit(db *sql.DB, couponId string, reduceBy int) error {
	var existingCount = 0
	row := db.QueryRow("SELECT userLimit from coupons where couponId = ?", couponId)
	err := row.Scan(&existingCount)
	if err != nil {
		return err
	}
	if existingCount < reduceBy {
		return errors.New("not enough user limit")
	}
	newUserLimit := existingCount - reduceBy
	_, err = db.Exec("UPDATE coupons set userLimit = ? where couponId = ?", newUserLimit, couponId)
	if err != nil {
		return err
	}
	return nil
}

func GetCouponDetails(db *sql.DB, coupon, email string) (*entities.CouponEntity, error) {
	row := db.QueryRow("SELECT userLimit,expiryDate,RuleSet from coupons where couponId = ? ", coupon)
	var couponDetail entities.CouponEntity
	var ruleInfo string
	err := row.Scan(&couponDetail.UserLimit, &couponDetail.ExpiryDate, &ruleInfo)
	if err != nil {
		return nil, err
	}
	couponDetail.CouponId = coupon
	currentDate := time.Now().UTC()
	if couponDetail.ExpiryDate.After(currentDate) && couponDetail.UserLimit > 0 {
		err := json.Unmarshal([]byte(ruleInfo), &couponDetail.Rule)
		if err != nil {
			return nil, err
		}
		return &couponDetail, nil
	} else {
		if couponDetail.UserLimit < 1 {
			fmt.Println("User limit reached")
		}
		if couponDetail.ExpiryDate.Before(currentDate) {
			fmt.Println("coupon has expired")
		}
		err := RemoveCouponFromCart(db, email)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("invalid coupon")
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
