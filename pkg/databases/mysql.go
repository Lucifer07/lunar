package databases

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlData struct {
	conn *gorm.DB
}

func (m *mysqlData) Connect(config DBConfig) error {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Database,
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}
	m.conn = db
	return nil
}
func (m *mysqlData) GetConnection() interface{} {
	return m.conn
}

