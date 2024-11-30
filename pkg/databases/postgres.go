package databases

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresData struct {
	Conn *gorm.DB
}

func (p *postgresData) Connect(config DBConfig) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	p.Conn = db
	return nil
}
func (p *postgresData) GetConnection() interface{} {
	return p.Conn
}