package model

type Error interface {
	Error() string
	Code() int
}

type NotEnoughFiat struct{}

func (e NotEnoughFiat) Error() string {
	return "Not Enough Fiat"
}

func (e NotEnoughFiat) Code() int {
	return 400
}

type NotEnoughBtc struct{}

func (e NotEnoughBtc) Error() string {
	return "Not Enough Btc"
}

func (e NotEnoughBtc) Code() int {
	return 400
}

type NotEnoughETH struct{}

func (e NotEnoughETH) Error() string {
	return "Not Enough ETH"
}

func (e NotEnoughETH) Code() int {
	return 400
}

type SqlInternalError struct{}

func (e SqlInternalError) Error() string {
	return "Sql Internal Error"
}

func (e SqlInternalError) Code() int {
	return 500
}

type NoRowsAffectedError struct{}

func (e NoRowsAffectedError) Error() string {
	return "Sql Internal Error"
}

func (e NoRowsAffectedError) Code() int {
	return 400
}
