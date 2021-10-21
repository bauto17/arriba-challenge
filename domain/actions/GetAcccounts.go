package actions

import "arriba-challenge/domain/model"

type GetAccounts struct {
	repo model.AccountRepository
}

func NewGetAccounts(repo model.AccountRepository) *GetAccounts {
	return &GetAccounts{
		repo: repo,
	}
}

func (ca *GetAccounts) Execute() ([]*model.Balance, error) {
	return ca.repo.GetAccounts()
}
