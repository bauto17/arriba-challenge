package actions

import "arriba-challenge/domain/model"

type Withdraw struct {
	repo model.AccountRepository
}

func NewWithdraw(repo model.AccountRepository) *Withdraw {
	return &Withdraw{
		repo: repo,
	}
}

func (ca *Withdraw) Execute(accountId int64, amount int64) (bool, model.Error) {
	return ca.repo.Withdraw(accountId, amount)
}
