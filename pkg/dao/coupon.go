package dao

import (
	"crypto/rand"
	"database/sql"
	"e-food/model"
	"encoding/base32"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type CouponHandler interface {
	InsertNewCoupon(db *sql.DB, userLimit int, expTime time.Time, ruleSet string) (string, error)
	ReduceUserLimit(db *sql.DB, couponId string, reduceBy int) error
	insertWithUniqueId(db *sql.DB, userLimit int, expTime time.Time, ruleSet, randId string) error
	generateRandomToken(length int) string
	GetCouponDetails(db *sql.DB, coupon string) (*model.CouponEntity, error)
}

type coupon struct{}

func CreateCouponHandler() CouponHandler {
	return &coupon{}
}

func (c *coupon) InsertNewCoupon(db *sql.DB, userLimit int, expTime time.Time, ruleSet string) (string, error) {
	randId := c.generateRandomToken(10)
	err := c.insertWithUniqueId(db, userLimit, expTime, ruleSet, randId)

	for err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		randId = c.generateRandomToken(10)
		err = c.insertWithUniqueId(db, userLimit, expTime, ruleSet, randId)
	}
	if err != nil {
		return "", err
	}
	return randId, nil
}

func (c *coupon) ReduceUserLimit(db *sql.DB, couponId string, reduceBy int) error {
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

func (c *coupon) insertWithUniqueId(db *sql.DB, userLimit int, expTime time.Time, ruleSet, randId string) error {
	_, err := db.Exec("INSERT INTO coupons (couponId, expiryDate, RuleSet, userLimit) VALUES (?,?,?,?)", randId, expTime, ruleSet, userLimit)
	if err != nil {
		return err
	}
	return nil
}

func (c *coupon) generateRandomToken(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}

func (c *coupon) GetCouponDetails(db *sql.DB, couponId string) (*model.CouponEntity, error) {
	row := db.QueryRow("SELECT userLimit,expiryDate,RuleSet from coupons where couponId = ? ", couponId)
	var couponDetail model.CouponEntity
	var ruleInfo string
	err := row.Scan(&couponDetail.UserLimit, &couponDetail.ExpiryDate, &ruleInfo)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(ruleInfo), &couponDetail.Rule)
	if err != nil {
		return nil, err
	}
	couponDetail.CouponId = couponId
	return &couponDetail, nil
}
