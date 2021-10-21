package model

type Balance struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Usd  int64  `json:"usd"`
	Btc  int64  `json:"btc"`
	Eth  int64  `json:"eth"`
}
