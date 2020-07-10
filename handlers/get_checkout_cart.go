package handlers

import (
	"database/sql"
	"e-food/pkg/dao"
	"e-food/pkg/rules"
	"e-food/pkg/utils"
	"e-food/restapi/operations/user"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/martian/log"
)

type cartPreviewImpl struct {
	dbClient *sql.DB
}

func NewCartCheckoutHandler(db *sql.DB) user.CheckoutHandler {
	return &cartPreviewImpl{
		dbClient: db,
	}
}

func (impl *cartPreviewImpl) Handle(params user.CheckoutParams, principal interface{}) middleware.Responder {
	email, err := utils.ValidateHeader(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		return user.NewCheckoutInternalServerError().WithPayload("error in parsing token")
	}
	cartItems, err := dao.GetCustomerCart(impl.dbClient, email.(string))
	ruleBooks, err := rules.LoadRuleBook()
	fmt.Println("===============")
	for _, v := range ruleBooks {
		fmt.Println(v.RuleId)
	}
	fmt.Println("===============")
	if err != nil {
		log.Errorf(err.Error())
		return user.NewCheckoutInternalServerError().WithPayload("Error getting info")
	}
	return user.NewCheckoutOK().WithPayload(cartItems)
}
