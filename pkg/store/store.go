package store

import (
	"context"
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	config	*Config
	db		*mongo.Client
}

func New() (IStore, error) {
	conf, err := loadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config, err: %v", err)
	}

	db := connectDB(conf)
	s := Store{config: conf, db: db}

	return s, nil
}

func connectDB(config *Config) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		fmt.Println(err)
	}
 
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	defer cancel()
 
	if err != nil {
		fmt.Println(err)
	}
 
	fmt.Println("Connected to mongoDB")
	return client
}

func loadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("STORE", &cfg); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &cfg, nil
}

func (s Store) Account() IAccountInterface {
	return accountStore{s}
}

func (s Store) Transaction() ITransactionInterface {
	return transactionStore{s}
}
