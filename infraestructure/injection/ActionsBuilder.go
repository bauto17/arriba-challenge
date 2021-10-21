package injection

import (
	"arriba-challenge/domain/actions"
	"arriba-challenge/infraestructure/config"
)

type Actions struct {
	*actions.CreateAccount
	*actions.Deposit
	*actions.Withdraw
	*actions.Buy
	*actions.Vender
	*actions.GetAccounts
}

func BuildActions(config config.GeneralConfig) (*Actions, error) {
	repositories, err := BuildRepositories(config)
	if err != nil {
		return nil, err
	}

	createAccountAction := actions.NewCreateAccount(repositories.AccountRepository)
	depositAction := actions.NewDeposit(repositories.AccountRepository)
	withdrawAction := actions.NewWithdraw(repositories.AccountRepository)
	buyAction := actions.NewBuy(repositories.AccountRepository, repositories.CurrencyRepository)
	venderAction := actions.NewVender(repositories.AccountRepository, repositories.CurrencyRepository)
	getAccunts := actions.NewGetAccounts(repositories.AccountRepository)

	return &Actions{
		CreateAccount: createAccountAction,
		Deposit:       depositAction,
		Withdraw:      withdrawAction,
		Buy:           buyAction,
		Vender:        venderAction,
		GetAccounts:   getAccunts,
	}, nil
}
