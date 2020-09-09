package expanders

import (
	"html/template"
	"math/big"

	acc "github.com/leekchan/accounting"
)

//All returns a list of all available expanders
func All() template.FuncMap {
	return template.FuncMap{
		"currency": currency,
	}
}

var accPerCountry = map[string]acc.Accounting{
	"US": acc.Accounting{Symbol: "$", Precision: 2},
	"DE": acc.Accounting{Symbol: "€", Precision: 2},
	"BR": acc.Accounting{Symbol: "R$ ", Precision: 2, Thousand: ".", Decimal: ","},
}

func currency(args ...interface{}) (string, error) {
	if len(args) == 1 {
		args = append([]interface{}{"US"}, args...)
	}

	ac := accPerCountry[args[0].(string)]
	return ac.FormatMoneyBigFloat(args[1].(*big.Float)), nil
}
