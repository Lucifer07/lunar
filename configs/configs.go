package configs

import (
	. "github.com/Lucifer07/lunar/pkg/databases"
	. "github.com/Lucifer07/lunar/utils"
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

var JwtSecreet = GetEnv("jwt_secret","secret")
