package integration

import (
	"e-food/models"
	"fmt"
)

func PrepareBilling(cartItems []*models.CartItem) (*models.BillableCart, error) {
	rules, err := CreateRuleBook()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	rules.ApplyRules(cartItems)
	return nil, nil
}
