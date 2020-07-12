package dao

import (
	"crypto/rand"
	"database/sql"
	"encoding/base32"
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
