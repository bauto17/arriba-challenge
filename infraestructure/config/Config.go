package config

import (
	"log"
	"os"
	"strconv"
)

type GeneralConfig struct {
	DatabaseConfig
	WebServerConfig
}

func GetConfig() GeneralConfig {
	return GeneralConfig{
		DatabaseConfig{
			Host:               getEnv("DB_HOST"),
			DbName:             getEnv("DB_NAME"),
			Password:           getEnv("DB_PASS"),
			User:               getEnv("DB_USER"),
			Port:               getIntEnv("DB_PORT"),
			PoolSize:           20,
			ConnLifeTimeSecond: 120,
		},
		WebServerConfig{
			Port: getEnv("PORT"),
		},
	}
}

type WebServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host               string
	Port               int
	DbName             string
	User               string
	Password           string
	PoolSize           int64
	ConnLifeTimeSecond int64
}

func getEnv(envName string) string {
	value := os.Getenv(envName)
	if value == "" {
		log.Fatalf("Wrong environment variable")
	}

	return value
}

func getIntEnv(envName string) int {
	value := getEnv(envName)
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Wrong environment variable type. Expected '%s' of type int", envName)
	}
	return result
}
