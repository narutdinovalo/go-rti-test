package domain

const (
	PriceTypeCost              = "COST"
	PriceTypeDiscount          = "Discount"
	OperatorEqual              = "EQ"
	OperatorGreaterThanOrEqual = "GTE"
	OperatorLessThanOrEqual    = "LTE"
)

type RuleApplicability struct {
	CodeName string `json:"codeName"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Price struct {
	Cost                float64             `json:"cost"`
	PriceType           string              `json:"priceType,omitempty"`
	RuleApplicabilities []RuleApplicability `json:"ruleApplicabilities,omitempty"`
}

type Component struct {
	Name   string  `json:"name"`
	IsMain bool    `json:"isMain,omitempty"`
	Prices []Price `json:"prices"`
}

type Product struct {
	Name       string      `json:"name"`
	Components []Component `json:"components"`
}

type Condition struct {
	RuleName string `json:"ruleName"`
	Value    string `json:"value"`
}

type Offer struct {
	Product
	TotalCost Price `json:"totalCost"`
}

var getProduct = Product{
	Name: "Игровой",
	Components: []Component{
		{
			IsMain: true,
			Name:   "Интернет",
			Prices: []Price{
				{
					Cost:      100,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "adsl",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "10",
						},
					},
				},
				{
					Cost:      150,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "adsl",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "15",
						},
					},
				},
				{
					Cost:      500,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "xpon",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "100",
						},
					},
				},
				{
					Cost:      900,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "xpon",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "200",
						},
					},
				},
				{
					Cost:      200,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "fttb",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "30",
						},
					},
				},
				{
					Cost:      400,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "fttb",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "50",
						},
					},
				},
				{
					Cost:      600,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "fttb",
						},
						{
							CodeName: "internetSpeed",
							Operator: OperatorEqual,
							Value:    "200",
						},
					},
				},
				{
					Cost:      10,
					PriceType: PriceTypeDiscount,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "internetSpeed",
							Operator: OperatorGreaterThanOrEqual,
							Value:    "50",
						},
					},
				},
				{
					Cost:      15,
					PriceType: PriceTypeDiscount,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "internetSpeed",
							Operator: OperatorGreaterThanOrEqual,
							Value:    "100",
						},
					},
				},
			},
		},
		{
			Name: "ADSL Модем",
			Prices: []Price{
				{
					Cost:      300,
					PriceType: PriceTypeCost,
					RuleApplicabilities: []RuleApplicability{
						{
							CodeName: "technology",
							Operator: OperatorEqual,
							Value:    "adsl",
						},
					},
				},
			},
		},
	},
}

func GetProduct() *Product {
	return &getProduct
}
