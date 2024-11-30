package databases

import (
	"context"

	"github.com/Lucifer07/lunar/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoData struct {
	conn *mongo.Client
}

func (m *mongoData) Connect(config DBConfig) error {
	clientOptions := options.Client().ApplyURI(config.URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return errors.ErrInvalidConfig
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return errors.ErrInvalidConnection
	}
	m.conn = client
	return nil
}
func (m *mongoData) GetConnection() interface{} {
	return m.conn
}
