package mongo

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	data *Data
	once sync.Once
)

type Data struct {
	DB *mongo.Client
}

func New() *Data {
	once.Do(initDB)
	return data
}

func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Fatal(err)
	}
	data = &Data{
		DB: db,
	}
}

func (data *Data) Ping() error {
	return data.DB.Ping(context.Background(), nil)
}

func (data *Data) Close() error {
	return data.DB.Disconnect(context.Background())
}
