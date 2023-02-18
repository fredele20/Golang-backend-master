package api

import (
	"github.com/fredele20/Golang-backend-master/util"
	"github.com/go-playground/validator/v10"
)


var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}