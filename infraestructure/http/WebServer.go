package http

import (
	"arriba-challenge/infraestructure/config"
	"arriba-challenge/infraestructure/http/handlers"
	"arriba-challenge/infraestructure/injection"
	"github.com/gin-gonic/gin"
)

type WebServer struct {
	actions *injection.Actions
	engine  *gin.Engine
	config  config.WebServerConfig
}

func NewWebServer(config config.WebServerConfig, actions *injection.Actions) *WebServer {
	accountHandler := handlers.NewAccountHandler(actions)
	r := gin.New()

	r.GET("", func(g *gin.Context) {
		g.String(200, "UP!")
	})
	// Simple group: v1
	v1 := r.Group("/account")
	{
		v1.POST("", accountHandler.CreateAccount)
		v1.GET("", accountHandler.GetAccounts)
		v1.POST("/:account_id/deposit", accountHandler.Deposit)
		v1.POST("/:account_id/withdraw", accountHandler.Withdraw)
		v1.POST("/:account_id/buy", accountHandler.Buy)
		v1.POST("/:account_id/vender", accountHandler.Vender)
	}

	return &WebServer{
		actions: actions,
		engine:  r,
		config:  config,
	}
}

func (ws *WebServer) Start() error {
	return ws.engine.Run(":" + ws.config.Port)
}
