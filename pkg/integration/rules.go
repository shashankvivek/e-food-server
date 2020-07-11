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

func (r *RuleCollection) ApplyRules(cartItems []*models.CartItem) error {
	//var offerCartItems []models.OfferItem
	//var nonOfferCartItems []models.BillingItem
	//ruleMap := make(map[string]*models.BillingItem)
	for _, rule := range r.RuleBook {
		productsFound := checkForMatchingProducts(rule.RuleSet, cartItems)
		if productsFound {
			//prodWithOfferApplied := extractProductsWithOffer(rule,cartItems)
			fmt.Println(rule.RuleId)
		} else {
			fmt.Println("Not matched: " + rule.RuleId)
		}
		//for _, filter := range rule.RuleSet {
		//	if filter.ProductId == item.ProductID {
		//		if *filter.MaxQuantity == *filter.MinQuantity {
		//			// exact match
		//		} else if filter.MinQuantity == nil && filter.MaxQuantity != nil {
		//			// max limit is set
		//		} else if filter.MinQuantity != nil && filter.MaxQuantity == nil {
		//			// min limit and no max limit
		//		} else if *filter.MinQuantity <= *filter.MaxQuantity {
		//			// there is a range in min and max qty
		//		} else {
		//			return errors.New("invalid rule: min can't be greater than max qty :" + rule.RuleId)
		//		}
		//	} else {
		//		filterApplicable = false
		//		break
		//	}
		//}
	}
	return nil
}

func extractProductsWithOffer(rule *Rule, cartItems []*models.CartItem) {

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
