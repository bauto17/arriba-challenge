package injection

import (
	"arriba-challenge/domain/model"
	"arriba-challenge/infraestructure/config"
	"arriba-challenge/infraestructure/repositories"
)

type Repositories struct {
	AccountRepository  model.AccountRepository
	CurrencyRepository model.CurrencyRepository
}

func BuildRepositories(config config.GeneralConfig) (*Repositories, error) {

	accountRepo, err := repositories.NewAccountRepository(config)
	if err != nil {
		return nil, err
	}

	currencyRepo := repositories.NewCurrencyRepository()

	return &Repositories{
		AccountRepository:  accountRepo,
		CurrencyRepository: currencyRepo,
	}, nil
}
