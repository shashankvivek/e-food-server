package entities

import (
	"e-food/models"
	"time"
)

type CouponEntity struct {
	CouponId   string
	ExpiryDate time.Time
	Rule       Rule
	UserLimit  int
}

type Rule struct {
	RuleId   string             `json:"ruleId"`
	Discount float64            `json:"discount,omitempty"`
	RuleSet  map[string]*Filter `json:"filters"`
}

type Filter struct {
	MinQuantity *int64 `json:"minQuantity,omitempty"`
	EqQunatity  *int64 `json:"eqQunatity,omitempty"`
}

type EngineTempItem struct {
	NumberOfGroupInSet     int
	TempItem               *models.CartItem
	ToBeReducedQtyFromCart int64
}
