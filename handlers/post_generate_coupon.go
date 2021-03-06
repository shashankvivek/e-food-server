package handlers

import (
	"database/sql"
	"e-food/api/models"
	"e-food/api/restapi/operations/admin"
	"e-food/model"
	"e-food/pkg/dao"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"time"
)

type generateCouponImpl struct {
	dbClient      *sql.DB
	couponHandler dao.CouponHandler
}

func NewAdminGenerateCouponHandler(db *sql.DB, couponHandler dao.CouponHandler) admin.GenerateCouponHandler {
	return &generateCouponImpl{
		dbClient:      db,
		couponHandler: couponHandler,
	}
}

func (impl *generateCouponImpl) Handle(params admin.GenerateCouponParams) middleware.Responder {
	defaultRuleSet := "{\"ruleId\": \"c1\",\"discount\": 30.00,\"filters\": {\"3\": {\"minQuantity\": 1}}}"
	expirationTime := time.Now().UTC().Add(90 * time.Second)
	userLimit := 1
	rule := model.Rule{}
	err := json.Unmarshal([]byte(defaultRuleSet), &rule)
	if err != nil {
		fmt.Println(err.Error())
		return admin.NewGenerateCouponInternalServerError().WithPayload("error decoding rule")
	}
	couponCode, err := impl.couponHandler.InsertNewCoupon(impl.dbClient, userLimit, expirationTime, defaultRuleSet)
	if err != nil {
		fmt.Println(err.Error())
		return admin.NewGenerateCouponInternalServerError().WithPayload("error generating coupon")
	}
	return admin.NewGenerateCouponOK().WithPayload(&models.SuccessResponse{Success: true, Message: couponCode})
}
