package settings

import (
	"github.com/Lucifer07/lunar/configs"
	. "github.com/Lucifer07/lunar/pkg/databases"
)

func InitDB() (Databases, error) {
	dbManager := NewDatabases()
	for _, db := range configs.Database {
		err := dbManager.AddDB(db)
		if err != nil {
			return nil, err
		}
	}
	return dbManager, nil
}

func InitRepo(name,connection string,dbManager Databases){
	db:= dbManager.GetDB()
	var database 
	switch connection {
	case "mongo":


}

