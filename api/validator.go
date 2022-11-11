package api

import (
	"github.com/go-playground/validator/v10"
	"simplebank/util"
)

var validCurrency validator.Func = func(feildLevel validator.FieldLevel) bool {
	if currency, ok := feildLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}

	return false
}
