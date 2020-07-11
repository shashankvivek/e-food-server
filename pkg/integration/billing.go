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
	offerItems, remainingItems, err := rules.ApplyRules(cartItems)
	var bItem []*models.BillingItem
	for _, v := range remainingItems {
		bItem = append(bItem, &models.BillingItem{
			Currency:    v.Currency,
			TotalPrice:  0,
			ProductID:   v.ProductID,
			Quantity:    v.Quantity,
			ProductName: v.ProductName,
			ImageURL:    v.ImageURL,
			UnitPrice:   v.UnitPrice,
		})
	}

	finalCart := &models.BillableCart{
		TotalPrice:  0,
		Currency:    "Rs",
		TotalSaving: 66,
		OfferItems:  offerItems,
		Items:       bItem,
	}

	return finalCart, nil
}
