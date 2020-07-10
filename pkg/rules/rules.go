package rules

import (
	"encoding/json"
	"io/ioutil"
)

type RuleBook struct {
	RuleId  string  `json:"ruleId"`
	RuleSet *[]Rule `json:"ruleSet"`
}

type Rule struct {
	Product     int64  `json:"productId"`
	MinQuantity int    `json:"minQuantity"`
	MaxQuantity int    `json:"maxQuantity"`
	Operator    string `json:"operator"`
}

func LoadRuleBook() ([]RuleBook, error) {
	data, err := ioutil.ReadFile("./rules/rules.json")
	if err != nil {
		return nil, err
	}
	var ruleBooks []RuleBook
	err = json.Unmarshal(data, &ruleBooks)
	if err != nil {
		return nil, err
	}
	return ruleBooks, nil
}
