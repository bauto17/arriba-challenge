package model

type AccountRepository interface {
	CreateAccount(name string) (*Balance, Error)
	GetAccounts() ([]*Balance, Error)
	Deposit(accountId int64, amount int64) (bool, Error)
	Withdraw(accountId int64, amount int64) (bool, Error)
	Buy(accountId int64, amount int64, currency string, currencyAmount int64) (bool, Error)
	Vender(accountId int64, amount int64, currency string, currencyAmount int64) (bool, Error)
}

type CurrencyRepository interface {
	GetCurrency(name string) (string, int64)
}
