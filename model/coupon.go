package model

import (
	"time"
)

type CouponEntity struct {
	CouponId   string
	ExpiryDate time.Time
	Rule       Rule
	UserLimit  int
}
