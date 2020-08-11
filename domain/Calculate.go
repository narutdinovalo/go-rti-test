package domain

import (
	"fmt"
	"go-rti-testing/infrastructure"
)

func Calculate(product *Product, conditions []Condition) (offer *Offer, err error) {
	infrastructure.Logger("Запущен расчет продуктового предложения")
	if product == nil && conditions == nil {
		return nil, nil
	}
	if product != nil && len(conditions) < 1 {
		return &Offer{}, nil
	}

	addComponent := []Component{}
	addProduct := Product{}
	totalCost := float64(0)
	isMain := bool(false)

	for _, component := range product.Components {
		addPrices := []Price{}
		addPrices, err = getPrice(component.Prices, conditions)
		if err != nil {
			return nil, err
		}
		if len(addPrices) > 0 {
			addComponent = append(addComponent, Component{
				Name:   component.Name,
				IsMain: component.IsMain,
				Prices: addPrices})
			if component.IsMain {
				isMain = true
			}
		}
		addProduct = Product{
			Name:       product.Name,
			Components: addComponent,
		}
		for _, cost := range addPrices {
			totalCost += cost.Cost
		}
		offer = &Offer{
			Product:   addProduct,
			TotalCost: Price{Cost: totalCost},
		}

	}
	infrastructure.Logger("Расчет продуктового предложения успешно выполнен")
	if !isMain {
		return nil, nil
	}
	return offer, nil
}

func getPrice(prices []Price, conditions []Condition) ([]Price, error) {
	infrastructure.Logger("Начат расчет цен")
	addPrice := []Price{}
	addDiscount := Price{}
	totalPrice := []Price{}

	for _, price := range prices {
		apply, err := applyRule(price.RuleApplicabilities, conditions)
		if err != nil {
			return totalPrice, err
		}
		if apply == true {
			switch price.PriceType {
			case PriceTypeDiscount:
				addDiscount = getDiscount(price, addDiscount)
			case PriceTypeCost:
				addPrice = append(addPrice, price)
				if len(addPrice) > 1 {
					return totalPrice, fmt.Errorf("У компонента может быть только 1 цена")
				}
			}
		}
	}
	for _, finalPrice := range addPrice {
		totalPrice = append(totalPrice, Price{
			Cost: applyDiscount(finalPrice.Cost, addDiscount.Cost),
		})
	}
	infrastructure.Logger("Расчет цен успешно выполнен")
	return totalPrice, nil
}

func getDiscount(newDiscount Price, discount Price) Price {
	if newDiscount.Cost > discount.Cost {
		return newDiscount
	}
	return discount
}

func applyDiscount(price float64, discount float64) float64 {
	totalPrice := float64(0)
	if discount != 0 {
		totalPrice = (100 - discount) / 100 * price
	} else {
		totalPrice = price
	}
	return totalPrice
}

func applyRule(rules []RuleApplicability, conditions []Condition) (apply bool, err error) {
	for _, rule := range rules {
		apply, err = applyCondition(rule, conditions)
		if err != nil {
			infrastructure.Logger("Применение условий к правилам применимости")
			return false, err
		}
		if apply == false {
			break
		}
	}
	return apply, nil
}

func applyCondition(rule RuleApplicability, conditions []Condition) (apply bool, err error) {
	for _, condition := range conditions {
		if rule.CodeName == condition.RuleName {
			switch rule.Operator {
			case OperatorEqual:
				{
					if rule.Value == condition.Value {
						apply = true
					} else {
						apply = false
					}
				}
			case OperatorGreaterThanOrEqual:
				{
					ruleVal, err := infrastructure.ConvStrToFloat(rule.Value)
					if err != nil {
						return false, err
					}
					conditionVal, err := infrastructure.ConvStrToFloat(condition.Value)
					if err != nil {
						return false, err
					}
					if conditionVal >= ruleVal {
						apply = true
					} else {
						apply = false
					}
				}
			case OperatorLessThanOrEqual:
				{
					ruleVal, err := infrastructure.ConvStrToFloat(rule.Value)
					if err != nil {
						return false, err
					}
					conditionVal, err := infrastructure.ConvStrToFloat(condition.Value)
					if err != nil {
						return false, err
					}
					if conditionVal <= ruleVal {
						apply = true
					} else {
						apply = false
					}
				}
			}
		}
	}
	return apply, nil
}
