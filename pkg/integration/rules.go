package integration

import (
	"e-food/api/models"
	"e-food/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type RuleCollection struct {
	RuleBook []model.Rule
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
	var ruleBook []model.Rule
	err = json.Unmarshal(data, &ruleBook)
	if err != nil {
		return nil, err
	}
	return &RuleCollection{
		RuleBook: ruleBook,
	}, nil
}

func (r *RuleCollection) AppendNewRules(newRule model.Rule) {
	r.RuleBook = append(r.RuleBook, newRule)
}

func (r *RuleCollection) ApplyRules(cartItems []*models.CartItem) ([]*models.OfferItem, []*models.CartItem, error) {
	var offerCartItems []*models.OfferItem
	for _, rule := range r.RuleBook {
		productsFound := CheckForMatchingProductsWithRuleSets(rule.RuleSet, cartItems)
		if productsFound {
			var offerItem []*models.OfferItem
			// this seems to be an expensive operation with existing nested code, can there be a better way ?
			offerItem, cartItems = extractProductsWithOffer(&rule, cartItems)
			offerCartItems = append(offerCartItems, offerItem...)
		} else {
			fmt.Println("Not matched: " + rule.RuleId)
		}
	}
	return offerCartItems, cartItems, nil
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
