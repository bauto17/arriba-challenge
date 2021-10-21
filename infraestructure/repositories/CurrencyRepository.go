package repositories

import "arriba-challenge/domain/model"

type Currency struct {
	name  string
	value int64
}

type CurrencyRepository struct {
	data map[string]Currency
}

func NewCurrencyRepository() model.CurrencyRepository {
	return &CurrencyRepository{
		data: map[string]Currency{
			"BTC": {
				name:  "btc",
				value: 25,
			},
			"ETH": {
				name:  "eth",
				value: 6,
			},
		},
	}
}

func (cr *CurrencyRepository) GetCurrency(name string) (string, int64) {
	item := cr.data[name]
	return item.name, item.value
}
