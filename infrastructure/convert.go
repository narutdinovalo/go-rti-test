package infrastructure

import (
	"fmt"
	"strconv"
)

func ConvStrToFloat(str string) (float64, error) {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("Невозможно строку %v конвиртировать в число %v", str, err)
	}
	return result, nil
}
