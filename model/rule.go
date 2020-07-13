package model

type Rule struct {
	RuleId   string             `json:"ruleId"`
	Discount float64            `json:"discount,omitempty"`
	RuleSet  map[string]*Filter `json:"filters"`
}

type Filter struct {
	MinQuantity *int64 `json:"minQuantity,omitempty"`
	EqQunatity  *int64 `json:"eqQunatity,omitempty"`
}
