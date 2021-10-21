package main

import (
	"arriba-challenge/infraestructure/config"
	"arriba-challenge/infraestructure/http"
	"arriba-challenge/infraestructure/injection"
)

func main() {
	cfg := config.GetConfig()
	actions, err := injection.BuildActions(cfg)
	if err != nil {
		return
	}
	server := http.NewWebServer(cfg.WebServerConfig, actions)
	err = server.Start()
	if err != nil {
		return
	}
}
