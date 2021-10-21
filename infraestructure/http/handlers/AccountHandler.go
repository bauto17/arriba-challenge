package handlers

import (
	"arriba-challenge/infraestructure/http/request"
	"arriba-challenge/infraestructure/injection"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	*injection.Actions
}

func NewAccountHandler(actions *injection.Actions) *AccountHandler {
	return &AccountHandler{
		Actions: actions,
	}
}

func (ah *AccountHandler) CreateAccount(c *gin.Context) {
	var json request.CreateAccountRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance, aErr := ah.Actions.CreateAccount.Execute(json.Name)
	if aErr != nil {
		c.String(aErr.Code(), aErr.Error())
		return
	}
	c.JSON(200, gin.H{"response": balance})
}

func (ah *AccountHandler) GetAccounts(c *gin.Context) {
	list, err := ah.Actions.GetAccounts.Execute()
	if err != nil {
		c.String(err.Code(), err.Error())
		return
	}
	c.JSON(200, gin.H{"response": list})
}

func (ah *AccountHandler) Deposit(c *gin.Context) {
	param := c.Param("account_id")
	accountId, err := strconv.ParseInt(param, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var json request.DepositRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance, aErr := ah.Actions.Deposit.Execute(accountId, json.Amount)
	if aErr != nil {
		c.String(aErr.Code(), aErr.Error())
	}
	c.JSON(200, gin.H{"response": balance})
}

func (ah *AccountHandler) Withdraw(c *gin.Context) {
	param := c.Param("account_id")
	accountId, err := strconv.ParseInt(param, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var json request.WithdrawRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance, aErr := ah.Actions.Withdraw.Execute(accountId, json.Amount)
	if aErr != nil {
		c.String(aErr.Code(), aErr.Error())
		return
	}
	c.JSON(200, gin.H{"response": balance})
}

func (ah *AccountHandler) Buy(c *gin.Context) {
	param := c.Param("account_id")
	accountId, err := strconv.ParseInt(param, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var json request.BuyRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance, aErr := ah.Actions.Buy.Execute(accountId, json.Currency, json.Amount)
	if aErr != nil {
		c.String(aErr.Code(), aErr.Error())
		return
	}
	c.JSON(200, gin.H{"response": balance})
}

func (ah *AccountHandler) Vender(c *gin.Context) {
	param := c.Param("account_id")
	accountId, err := strconv.ParseInt(param, 0, 64)

	var json request.VenderRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance, aErr := ah.Actions.Vender.Execute(accountId, json.Currency, json.Amount)
	if aErr != nil {
		c.String(aErr.Code(), err.Error())
		return
	}
	c.JSON(200, gin.H{"response": balance})
}
