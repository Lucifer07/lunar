package main

import (
	"fmt"
	"log"

	"github.com/Lucifer07/lunar/pkg/databases"
	"gorm.io/gorm"
)

func main() {
	dbManager := databases.NewDatabases()
	dbConfig := databases.DBConfig{
		Name:     "mysql",
		Type:     "mysql",
		Host:     "localhost",
		Port:     "3306",
		User:     "test",
		Password: "test",
		Database: "test",
	}
	err := dbManager.AddDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to add database: %v", err)
	} else {
		fmt.Println("Database added successfully!")
	}
	db := dbManager.GetDB()
	mysqlDB := db["mysql"].GetConnection().(*gorm.DB)
	fmt.Println(mysqlDB)

}
