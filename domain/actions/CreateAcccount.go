package actions

import "arriba-challenge/domain/model"

type CreateAccount struct {
	repo model.AccountRepository
}

func NewCreateAccount(repo model.AccountRepository) *CreateAccount {
	return &CreateAccount{
		repo: repo,
	}
}

func (ca *CreateAccount) Execute(name string) (*model.Balance, model.Error) {
	return ca.repo.CreateAccount(name)
}
