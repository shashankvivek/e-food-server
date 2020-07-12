package integration

import (
	"e-food/models"
	"e-food/pkg/entities"
	"fmt"
)

func PrepareBilling(cartItems []*models.CartItem, couponInfo *entities.CouponEntity) (*models.BillableCart, error) {
	rules, err := CreateRuleBook()
	currencyVal := cartItems[0].Currency
	couponId := ""
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if couponInfo != nil {
		rules.AppendNewRules(couponInfo.Rule)
		couponId = couponInfo.CouponId
	}

	offerItems, remainingItems, _ := rules.ApplyRules(cartItems)
	var nonOfferItems []*models.BillingItem
	for _, v := range remainingItems {
		nonOfferItems = append(nonOfferItems, &models.BillingItem{
			Currency:    v.Currency,
			TotalPrice:  float64(v.Quantity) * v.UnitPrice,
			ProductID:   v.ProductID,
			Quantity:    v.Quantity,
			ProductName: v.ProductName,
			ImageURL:    v.ImageURL,
			UnitPrice:   v.UnitPrice,
		})
	}

	totalCartPrice, totalSavedAmount := getTotalCartPrice(offerItems, nonOfferItems)
	finalCart := &models.BillableCart{
		TotalPrice:  totalCartPrice,
		Currency:    currencyVal,
		TotalSaving: totalSavedAmount,
		OfferItems:  offerItems,
		Items:       nonOfferItems,
		CouponID:    couponId,
	}

	return finalCart, nil
}

func getTotalCartPrice(OfferItems []*models.OfferItem, nonOfferItems []*models.BillingItem) (float64, float64) {
	totalPrice := 0.0
	totalSaving := 0.0
	for _, item := range OfferItems {
		totalPrice = totalPrice + item.DiscountedPrice
		totalSaving = totalSaving + item.ActualPrice - item.DiscountedPrice
	}
	for _, item := range nonOfferItems {
		totalPrice = totalPrice + item.TotalPrice
	}

	return totalPrice, totalSaving
}
