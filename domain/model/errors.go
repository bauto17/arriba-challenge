package model

type NotEnoughFiat struct{}

func (e NotEnoughFiat) Error() string {
	return "Not Enough Fiat"
}

type NotEnoughBtc struct{}

func (e NotEnoughBtc) Error() string {
	return "Not Enough Btc"
}

type NotEnoughETH struct{}

func (e NotEnoughETH) Error() string {
	return "Not Enough ETH"
}

type SqlInternalError struct{}

func (e SqlInternalError) Error() string {
	return "Sql Internal Error"
}
