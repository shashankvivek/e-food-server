package integration

import (
	"e-food/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type RuleCollection struct {
	RuleBook []Rule
}

type Rule struct {
	RuleId   string             `json:"ruleId"`
	Discount float64            `json:"discount,omitempty"`
	RuleSet  map[string]*Filter `json:"filters"`
}

type Filter struct {
	MinQuantity *int64 `json:"minQuantity"`
	MaxQuantity *int64 `json:"maxQuantity"`
}

type EngineTempItem struct {
	NumberOfGroupInSet     int
	TempItem               *models.CartItem
	ToBeReducedQtyFromCart int64
}

func CreateRuleBook() (*RuleCollection, error) {
	// assumption that min and max qty are checked by system before being consumed here
	data, err := ioutil.ReadFile("./pkg/integration/rules.json")
	if err != nil {
		return nil, err
	}
	var ruleBook []Rule
	err = json.Unmarshal(data, &ruleBook)
	if err != nil {
		return nil, err
	}
	return &RuleCollection{
		RuleBook: ruleBook,
	}, nil
}

func (r *RuleCollection) ApplyRules(cartItems []*models.CartItem) ([]*models.OfferItem, []*models.CartItem, error) {
	var offerCartItems []*models.OfferItem
	//var nonOfferCartItems []models.BillingItem
	for _, rule := range r.RuleBook {
		productsFound := checkForMatchingProducts(rule.RuleSet, cartItems)
		if productsFound {
			fmt.Println("=====================================")
			fmt.Println(rule.RuleId)
			var offerItem []*models.OfferItem
			offerItem, cartItems = extractProductsWithOffer(&rule, cartItems)
			offerCartItems = append(offerCartItems, offerItem...)
		} else {
			fmt.Println("Not matched: " + rule.RuleId)
		}
	}
	fmt.Println(offerCartItems)
	fmt.Println(cartItems)
	return offerCartItems, cartItems, nil
}

func extractProductsWithOffer(rule *Rule, cartItems []*models.CartItem) ([]*models.OfferItem, []*models.CartItem) {
	var remainingCartItems []*models.CartItem
	var eligibleItems []*models.CartItem
	var offering []*models.OfferItem
	for _, item := range cartItems {
		if rule.RuleSet[strconv.FormatInt(item.ProductID, 10)] != nil {
			eligibleItems = append(eligibleItems, item)
		} else {
			remainingCartItems = append(remainingCartItems, item)
		}
	}
	maxSetPossible := int64(999999999)
	for _, product := range eligibleItems {
		filterRule := rule.RuleSet[strconv.FormatInt(product.ProductID, 10)]
		// min limit and no max limit
		if filterRule.MinQuantity != nil &&
			filterRule.MaxQuantity == nil &&
			product.Quantity >= *filterRule.MinQuantity {
			offering = append(offering, &models.OfferItem{
				RuleSetID: rule.RuleId,
				Items: []*models.BillingItem{
					{
						ProductID:   product.ProductID,
						Quantity:    product.Quantity,
						Currency:    product.Currency,
						UnitPrice:   product.UnitPrice,
						ImageURL:    product.ImageURL,
						ProductName: product.ProductName,
						TotalPrice:  product.UnitPrice * float64(product.Quantity),
					},
				},
			})
			indexInMainCart := getItemIndexInCart(*product, cartItems)
			return offering, append(cartItems[:indexInMainCart], cartItems[indexInMainCart+1:]...)
			//remove this entire item from cartItems
		} else if *filterRule.MaxQuantity == *filterRule.MinQuantity {
			//exact limit for rule
			setPossible := product.Quantity / *filterRule.MinQuantity
			if maxSetPossible > setPossible {
				maxSetPossible = setPossible
			}
		}

	}
	//loopCount := int64(0)
	for 0 < maxSetPossible {
		var items []*models.BillingItem
		for i, product := range eligibleItems {
			qtyAsPerFilter := *rule.RuleSet[strconv.FormatInt(product.ProductID, 10)].MinQuantity
			items = append(items, &models.BillingItem{
				ProductID:   product.ProductID,
				Quantity:    qtyAsPerFilter,
				Currency:    product.Currency,
				UnitPrice:   product.UnitPrice,
				ImageURL:    product.ImageURL,
				ProductName: product.ProductName,
				TotalPrice:  product.UnitPrice * float64(qtyAsPerFilter),
			})
			product.Quantity = product.Quantity - *rule.RuleSet[strconv.FormatInt(product.ProductID, 10)].MinQuantity
			if product.Quantity == 0 {
				eligibleItems = append(eligibleItems[:i], eligibleItems[i+1:]...)
			}
		}
		offering = append(offering, &models.OfferItem{
			RuleSetID: rule.RuleId,
			Items:     items,
		})
		maxSetPossible--
	}
	leftOverItems := eligibleItems
	remainingCartItems = append(remainingCartItems, leftOverItems...)

	return offering, remainingCartItems

	//if *filter.MaxQuantity == *filter.MinQuantity {
	//	// exact match
	//} else if filter.MinQuantity == nil && filter.MaxQuantity != nil {
	//	// max limit is set
	//} else if filter.MinQuantity != nil && filter.MaxQuantity == nil {
	//	// min limit and no max limit
	//} else if *filter.MinQuantity <= *filter.MaxQuantity {
	//	// there is a range in min and max qty
	//} else {
	//	return errors.New("invalid rule: min can't be greater than max qty :" + rule.RuleId)
	//}
}

func getItemIndexInCart(item models.CartItem, items []*models.CartItem) int {
	for i, v := range items {
		if v.ProductID == item.ProductID {
			return i
		}
	}
	return -1
}
func checkForMatchingProducts(ruleSet map[string]*Filter, cartItems []*models.CartItem) bool {
	found := false
	matchedProdCount := 0
	for _, item := range cartItems {
		filterParams := ruleSet[strconv.FormatInt(item.ProductID, 10)]
		found = filterParams != nil
		if found && (filterParams.MinQuantity == nil || item.Quantity >= *filterParams.MinQuantity) {
			matchedProdCount++
		}
	}
	return matchedProdCount == len(ruleSet)
}

func exactMatchRule() {}

func maxLimitRule() {}

func minLimitRule() {}

func rangeLimitRule() {}
