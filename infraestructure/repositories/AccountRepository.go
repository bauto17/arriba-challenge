package repositories

import (
	"arriba-challenge/domain/model"
	"arriba-challenge/infraestructure/config"
	"context"
	"strings"
)

type AccountRepository struct {
	db Database
}

func NewAccountRepository(config config.GeneralConfig) (model.AccountRepository, error) {
	db, err := NewDatabase(config.DatabaseConfig)
	if err != nil {
		return nil, err
	}

	return &AccountRepository{
		db: db,
	}, nil
}

func (ar *AccountRepository) CreateAccount(name string) (*model.Balance, model.Error) {
	var response model.Balance
	row := ar.db.QueryRow(
		context.Background(),
		CreateAccountQuery,
		name,
	)
	err := row.Scan(&response.Id)
	response.Name = name

	if err != nil {
		return nil, Handle(err)
	}

	return &response, nil
}

func (ar *AccountRepository) GetAccounts() ([]*model.Balance, model.Error) {
	var response []*model.Balance
	rows, err := ar.db.Query(
		context.Background(),
		GetAccountsQuery,
	)
	if err != nil {
		return nil, Handle(err)
	}

	var scanError error
	for rows.Next() {
		item := model.Balance{}
		scanError = rows.Scan(
			&item.Id,
			&item.Name,
			&item.Usd,
			&item.Btc,
			&item.Eth,
		)

		if scanError == nil {
			response = append(response, &item)
		} else {
			if scanError != nil {
				return nil, Handle(scanError)
			}
		}
	}

	return response, nil
}

func (ar *AccountRepository) Deposit(accountId int64, amount int64) (bool, model.Error) {
	row, err := ar.db.Exec(
		context.Background(),
		DepositQuery,
		amount,
		accountId,
	)
	if err != nil {
		return false, Handle(err)
	}

	rows := row.RowsAffected()
	if rows == 0 {
		return false, model.NoRowsAffectedError{}
	}

	return true, nil
}

func (ar *AccountRepository) Withdraw(accountId int64, amount int64) (bool, model.Error) {
	row, err := ar.db.Exec(
		context.Background(),
		WithdrawQuery,
		amount,
		accountId,
	)
	if err != nil {
		return false, Handle(err)
	}

	rows := row.RowsAffected()
	if rows == 0 {
		return false, model.NoRowsAffectedError{}
	}

	return true, nil
}

func (ar *AccountRepository) Buy(accountId int64, amount int64, currency string, currencyAmount int64) (bool, model.Error) {
	row, err := ar.db.Exec(
		context.Background(),
		strings.Replace(BuyQuery, ":currency", currency, -1),
		amount,
		currencyAmount,
		accountId,
	)
	if err != nil {
		return false, Handle(err)
	}

	rows := row.RowsAffected()
	if rows == 0 {
		return false, model.NoRowsAffectedError{}
	}

	return true, nil
}

func (ar *AccountRepository) Vender(accountId int64, amount int64, currency string, currencyAmount int64) (bool, model.Error) {
	row, err := ar.db.Exec(
		context.Background(),
		strings.Replace(VenderQuery, ":currency", currency, -1),
		amount,
		currencyAmount,
		accountId,
	)

	if err != nil {
		return false, Handle(err)
	}

	rows := row.RowsAffected()
	if rows == 0 {
		return false, model.NoRowsAffectedError{}
	}

	return true, nil
}
