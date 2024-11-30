package databases

import "github.com/Lucifer07/lunar/pkg/errors"


type database interface {
	Connect(config DBConfig) error
	GetConnection() interface{}
}
type Databases interface {
	AddDB(config DBConfig) error
	GetDB() map[string]database
}
type databases struct {
	DB map[string]database
}

func NewDatabases() Databases {
	return &databases{
		DB: make(map[string]database),
	}
}
func (d *databases) AddDB(config DBConfig) error {
	var db database
	switch config.Type {
	case "mysql":
		db = &mysqlData{}
	case "postgres":
		db = &postgresData{}
	case "mongo":
		db = &mongoData{}
	default:
		return errors.ErrInvalidConnection
	}
	if err := db.Connect(config); err != nil {
		return errors.ErrInvalidConfig
	}
	d.DB[config.Name] = db
	return nil
}
func (d *databases) GetDB() map[string]database {
	return d.DB
}
