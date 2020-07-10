package integration

import (
	"e-food/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type RuleCollection struct {
	RuleBook []Rule
}

type Rule struct {
	RuleId  string    `json:"ruleId"`
	RuleSet *[]Filter `json:"filters"`
}

type Filter struct {
	Product     int64  `json:"productId"`
	MinQuantity int    `json:"minQuantity"`
	MaxQuantity int    `json:"maxQuantity"`
	Operator    string `json:"operator"`
}

func CreateRuleBook() (*RuleCollection, error) {
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

func (r *RuleCollection) ApplyRules(cartItems []*models.CartItem) {
	fmt.Println("===============")
	for _, v := range r.RuleBook {
		fmt.Println(v.RuleId)
	}
	fmt.Println("===============")
	//return nil
}
