package actions

import "arriba-challenge/domain/model"

type Vender struct {
	repo         model.AccountRepository
	currencyRepo model.CurrencyRepository
}

func NewVender(repo model.AccountRepository, currencyRepo model.CurrencyRepository) *Vender {
	return &Vender{
		repo:         repo,
		currencyRepo: currencyRepo,
	}
}

func (ca *Vender) Execute(accountId int64, currency string, amount int64) (bool, error) {
	idC, v := ca.currencyRepo.GetCurrency(currency)

	return ca.repo.Vender(accountId, v*amount, idC, amount)
}
