package configs

import (
	. "github.com/Lucifer07/lunar/pkg/databases"
)

var Database = []DBConfig{
	{
		Name:     "FirstDB",
		Type:     GetEnv("DB_CONNECTION", "mysql"),
		Host:     GetEnv("DB_HOST", "localhost"),
		Port:     GetEnv("DB_PORT", "3306"),
		User:     GetEnv("DB_USER", "root"),
		Password: GetEnv("DB_PASSWORD", ""),
		Database: GetEnv("DB_DATABASE", "test"),
	},
}
var JwtSecret = GetEnv("JWT_SECRET","secret")
var AppMode = GetEnv("APP_MODE","development")
var AppPort = GetEnv("APP_PORT","8080")
var AppHost = GetEnv("APP_HOST","localhost")
var ApiLocal = GetEnv("API_LOCAL","http://localhost:8080")
var ApiProd = GetEnv("API_PROD","http://localhost:8080")