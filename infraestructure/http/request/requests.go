package request

type CreateAccountRequest struct {
	Name string `json:"name" binding:"required"`
}

type DepositRequest struct {
	Amount int64 `json:"amount" binding:"required"`
}

type WithdrawRequest struct {
	Amount int64 `json:"amount" binding:"required"`
}

type BuyRequest struct {
	Currency string `json:"currency" binding:"required"`
	Amount   int64  `json:"amount" binding:"required"`
}

type VenderRequest struct {
	Currency string `json:"currency" binding:"required"`
	Amount   int64  `json:"amount" binding:"required"`
}
