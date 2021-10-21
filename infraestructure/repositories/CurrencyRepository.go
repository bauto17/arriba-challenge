package repositories

import (
	"arriba-challenge/domain/model"
	"arriba-challenge/infraestructure/config"
)

type Currency struct {
	name  string
	value int
}

type CurrencyRepository struct {
	data map[string]Currency
}

func NewCurrencyRepository(config config.CurrencyConfig) model.CurrencyRepository {
	return &CurrencyRepository{
		data: map[string]Currency{
			"BTC": {
				name:  "btc",
				value: config.Btc,
			},
			"ETH": {
				name:  "eth",
				value: config.Eth,
			},
		},
	}
}

func (cr *CurrencyRepository) GetCurrency(name string) (string, int64) {
	item := cr.data[name]
	return item.name, int64(item.value)
}
