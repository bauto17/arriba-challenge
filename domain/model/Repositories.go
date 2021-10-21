package model

type AccountRepository interface {
	CreateAccount(name string) (*Balance, error)
	GetAccounts() ([]*Balance, error)
	Deposit(accountId int64, amount int64) (bool, error)
	Withdraw(accountId int64, amount int64) (bool, error)
	Buy(accountId int64, amount int64, currency string, currencyAmount int64) (bool, error)
	Vender(accountId int64, amount int64, currency string, currencyAmount int64) (bool, error)
}

type CurrencyRepository interface {
	GetCurrency(name string) (string, int64)
}
