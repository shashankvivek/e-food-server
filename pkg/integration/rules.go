package integration

import (
	"e-food/api/models"
	"e-food/pkg/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type RuleCollection struct {
	RuleBook []entities.Rule
}

/*
Note :
1. rules.json will be created based on some rule-system exposed to sellers. It'll not be in DB.
2. The rule-system will be responsible for maintaining the valid struct of "rules.json"
3. The "filters" property has "ProductId" as "key"
4. The "ProductId" will contain "minQuantity" or  "eqQunatity", which can be combined to create different scenarios.
*/

func CreateRuleBook() (*RuleCollection, error) {

	data, err := ioutil.ReadFile("./resources/rules.json")
	if err != nil {
		return nil, err
	}
	var ruleBook []entities.Rule
	err = json.Unmarshal(data, &ruleBook)
	if err != nil {
		return nil, err
	}
	return &RuleCollection{
		RuleBook: ruleBook,
	}, nil
}

func (r *RuleCollection) AppendNewRules(newRule entities.Rule) {
	r.RuleBook = append(r.RuleBook, newRule)
}

func (r *RuleCollection) ApplyRules(cartItems []*models.CartItem) ([]*models.OfferItem, []*models.CartItem, error) {
	var offerCartItems []*models.OfferItem
	for _, rule := range r.RuleBook {
		productsFound := CheckForMatchingProductsWithRuleSets(rule.RuleSet, cartItems)
		if productsFound {
			var offerItem []*models.OfferItem
			offerItem, cartItems = extractProductsWithOffer(&rule, cartItems)
			offerCartItems = append(offerCartItems, offerItem...)
		} else {
			fmt.Println("Not matched: " + rule.RuleId)
		}
	}
	return offerCartItems, cartItems, nil
}

func extractProductsWithOffer(rule *entities.Rule, cartItems []*models.CartItem) ([]*models.OfferItem, []*models.CartItem) {
	var remainingCartItems []*models.CartItem
	var eligibleItems []*models.CartItem
	var offering []*models.OfferItem
	var leftOverItems []*models.CartItem
	for _, item := range cartItems {
		if rule.RuleSet[strconv.FormatInt(item.ProductID, 10)] != nil {
			eligibleItems = append(eligibleItems, item)
		} else {
			remainingCartItems = append(remainingCartItems, item)
		}
	}
	maxSetPossible := int64(999999999)
	hasEqualQuantityRule := false
	hasMinQuantityRule := false
	for _, product := range eligibleItems {
		filterRule := rule.RuleSet[strconv.FormatInt(product.ProductID, 10)]
		//exact limit for rule
		if filterRule.EqQunatity != nil {
			hasEqualQuantityRule = true
			setPossible := product.Quantity / *filterRule.EqQunatity
			if maxSetPossible > setPossible {
				maxSetPossible = setPossible
			}
		} else if filterRule.MinQuantity != nil &&
			product.Quantity >= *filterRule.MinQuantity {
			hasMinQuantityRule = true
			setPossible := int64(0)
			if product.Quantity >= *filterRule.MinQuantity {
				setPossible = 1
			}
			if maxSetPossible > setPossible {
				maxSetPossible = setPossible
			}
		}
	}
	if hasMinQuantityRule && hasEqualQuantityRule {
		leftOverItems, offering = WithBothRule(rule, eligibleItems)
	} else if hasEqualQuantityRule {
		leftOverItems, offering = groupItemsByOfferWithEqualQtyRule(rule, eligibleItems, maxSetPossible)
	} else if hasMinQuantityRule {
		leftOverItems, offering = withJustMinValue(rule, eligibleItems)
	}

	remainingCartItems = append(remainingCartItems, leftOverItems...)

	return offering, remainingCartItems
}

func WithBothRule(rule *entities.Rule, eligibleItems []*models.CartItem) ([]*models.CartItem, []*models.OfferItem) {
	var offering []*models.BillingItem
	var leftOverItems []*models.CartItem
	actualPrice := float64(0)
	for _, product := range eligibleItems {
		eqQtyAsPerFilter := rule.RuleSet[strconv.FormatInt(product.ProductID, 10)].EqQunatity
		minQtyAsPerFilterRef := rule.RuleSet[strconv.FormatInt(product.ProductID, 10)].MinQuantity
		if minQtyAsPerFilterRef != nil {
			// push all items as offer item
			offering = append(offering, &models.BillingItem{
				ProductID:   product.ProductID,
				Quantity:    product.Quantity,
				Currency:    product.Currency,
				UnitPrice:   product.UnitPrice,
				ImageURL:    product.ImageURL,
				ProductName: product.ProductName,
				TotalPrice:  product.UnitPrice * float64(product.Quantity),
			})
			product.Quantity = 0
			actualPrice = actualPrice + product.UnitPrice*float64(product.Quantity)
		} else {
			offering = append(offering, &models.BillingItem{
				ProductID:   product.ProductID,
				Quantity:    *eqQtyAsPerFilter,
				Currency:    product.Currency,
				UnitPrice:   product.UnitPrice,
				ImageURL:    product.ImageURL,
				ProductName: product.ProductName,
				TotalPrice:  product.UnitPrice * float64(product.Quantity),
			})
			product.Quantity = product.Quantity - *eqQtyAsPerFilter
			actualPrice = actualPrice + product.UnitPrice*float64(*eqQtyAsPerFilter)
		}
	}

	for _, prod := range eligibleItems {
		if prod.Quantity != 0 {
			leftOverItems = append(leftOverItems, prod)
		}
	}
	return leftOverItems, []*models.OfferItem{
		{
			ActualPrice:     actualPrice,
			RuleSetID:       rule.RuleId,
			DiscountedPrice: actualPrice * (1 - (rule.Discount / 100)),
			DiscountPercent: rule.Discount,
			Items:           offering,
		},
	}

}

func withJustMinValue(rule *entities.Rule, eligibleItems []*models.CartItem) ([]*models.CartItem, []*models.OfferItem) {
	var items []*models.BillingItem
	actualPrice := float64(0)

	// there is no possibility of sets creation
	for _, product := range eligibleItems {
		items = append(items, &models.BillingItem{
			ProductID:   product.ProductID,
			Quantity:    product.Quantity,
			Currency:    product.Currency,
			UnitPrice:   product.UnitPrice,
			ImageURL:    product.ImageURL,
			ProductName: product.ProductName,
			TotalPrice:  product.UnitPrice * float64(product.Quantity),
		})
		actualPrice = actualPrice + product.UnitPrice*float64(product.Quantity)
	}
	return nil, []*models.OfferItem{
		{
			RuleSetID:       rule.RuleId,
			Items:           items,
			ActualPrice:     actualPrice,
			DiscountPercent: rule.Discount,
			DiscountedPrice: actualPrice * (1 - (rule.Discount / 100)),
		},
	}
}

func groupItemsByOfferWithEqualQtyRule(rule *entities.Rule, eligibleItems []*models.CartItem, maxSetPossible int64) ([]*models.CartItem, []*models.OfferItem) {
	var leftOverItems []*models.CartItem
	var offering []*models.OfferItem
	for 0 < maxSetPossible {
		totalOfferItemPrice := 0.0
		var items []*models.BillingItem
		for _, product := range eligibleItems {
			qtyAsPerFilter := *rule.RuleSet[strconv.FormatInt(product.ProductID, 10)].EqQunatity
			items = append(items, &models.BillingItem{
				ProductID:   product.ProductID,
				Quantity:    qtyAsPerFilter,
				Currency:    product.Currency,
				UnitPrice:   product.UnitPrice,
				ImageURL:    product.ImageURL,
				ProductName: product.ProductName,
				TotalPrice:  product.UnitPrice * float64(qtyAsPerFilter),
			})
			product.Quantity = product.Quantity - qtyAsPerFilter
			totalOfferItemPrice = totalOfferItemPrice + product.UnitPrice*float64(qtyAsPerFilter)
		}
		offering = append(offering, &models.OfferItem{
			RuleSetID:       rule.RuleId,
			ActualPrice:     totalOfferItemPrice,
			DiscountPercent: rule.Discount,
			DiscountedPrice: totalOfferItemPrice * (1 - (rule.Discount / 100)),
			Items:           items,
		})
		maxSetPossible--
	}
	for _, prod := range eligibleItems {
		if prod.Quantity != 0 {
			leftOverItems = append(leftOverItems, prod)
		}
	}
	return leftOverItems, offering
}

func CheckForMatchingProductsWithRuleSets(ruleSet map[string]*entities.Filter, cartItems []*models.CartItem) bool {
	found := false
	matchedProdCount := 0
	for _, item := range cartItems {
		filterParams := ruleSet[strconv.FormatInt(item.ProductID, 10)]
		found = filterParams != nil
		if found &&
			((filterParams.MinQuantity != nil && item.Quantity >= *filterParams.MinQuantity) ||
				(filterParams.EqQunatity != nil && item.Quantity >= *filterParams.EqQunatity)) {
			matchedProdCount++
		}

	}
	return matchedProdCount == len(ruleSet)
}

/**
[
  {
    "ruleId": "r1",
    "discount": 10.00,
    "filters": {
      "1": {
        "minQuantity": 7
      }
    }
  },
  {
    "ruleId": "r2",
    "discount": 30.00,
    "filters": {
      "2": {
        "minQuantity": 5
      },
      "3": {
        "eqQunatity": 2
      },
      "5": {
        "eqQunatity": 4
      }
    }
  },
  {
    "ruleId": "r3",
    "discount": 10.00,
    "filters": {
      "4": {
        "eqQunatity": 4
      }
    }
  }
]
*/
