package handlers

import (
	"database/sql"
	"e-food/pkg/integration"
	"e-food/restapi/operations/admin"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
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
	//expirationTime := time.Now().Add(10* time.Second)

	rule := &integration.Rule{}

	err := json.Unmarshal([]byte(defaultRuleSet), rule)
	if err != nil {
		return admin.NewGenerateCouponInternalServerError().WithPayload("error decoding rule")
	}
	fmt.Println(*rule.RuleSet["4"].MinQuantity)

	return nil
}
