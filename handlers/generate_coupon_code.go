package handlers

import (
	"database/sql"
	"e-food/models"
	"e-food/pkg/dao"
	"e-food/pkg/integration"
	"e-food/restapi/operations/admin"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"time"
)

type generateCouponImpl struct {
	dbClient *sql.DB
}

func NewAdminGenerateCouponHandler(db *sql.DB) admin.GenerateCouponHandler {
	return &generateCouponImpl{
		dbClient: db,
	}
}

func (impl *generateCouponImpl) Handle(params admin.GenerateCouponParams) middleware.Responder {
	defaultRuleSet := "{\"ruleId\": \"c1\",\"discount\": 30.00,\"filters\": {\"4\": {\"minQunatity\": 1}}}"
	expirationTime := time.Now().UTC().Add(900 * time.Second)
	userLimit := 1
	rule := &integration.Rule{}
	err := json.Unmarshal([]byte(defaultRuleSet), rule)
	if err != nil {
		return admin.NewGenerateCouponInternalServerError().WithPayload("error decoding rule")
	}
	couponCode, err := dao.InsertNewCoupon(impl.dbClient, userLimit, expirationTime, defaultRuleSet)
	if err != nil {
		fmt.Println(err.Error())
		return admin.NewGenerateCouponInternalServerError().WithPayload("error generating coupon")
	}
	return admin.NewGenerateCouponOK().WithPayload(&models.SuccessResponse{Success: true, Message: couponCode})
}
