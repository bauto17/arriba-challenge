package actions

import "arriba-challenge/domain/model"

type Buy struct {
	repo         model.AccountRepository
	currencyRepo model.CurrencyRepository
}

func NewBuy(repo model.AccountRepository, currencyRepo model.CurrencyRepository) *Buy {
	return &Buy{
		repo:         repo,
		currencyRepo: currencyRepo,
	}
}

func (ca *Buy) Execute(accountId int64, currency string, amount int64) (bool, model.Error) {
	idC, v := ca.currencyRepo.GetCurrency(currency)

	return ca.repo.Buy(accountId, v*amount, idC, amount)
}
