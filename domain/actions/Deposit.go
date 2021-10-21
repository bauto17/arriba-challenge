package actions

import "arriba-challenge/domain/model"

type Deposit struct {
	repo model.AccountRepository
}

func NewDeposit(repo model.AccountRepository) *Deposit {
	return &Deposit{
		repo: repo,
	}
}

func (ca *Deposit) Execute(accountId int64, amount int64) (bool, model.Error) {
	return ca.repo.Deposit(accountId, amount)
}
