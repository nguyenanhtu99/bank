package store

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	validatorAccount = bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"username", "balance"},
			"properties": bson.M{
				"username": bson.M{
					"bsonType": "string",
					"description": "Must be a string",
				},
				"balance": bson.M{
					"bsonType": "double",
					"minimum": 0,
					"description": "Must be greater than 0",
				},
			},
		},
	}
	validatorTransaction = bson.M{
		"$jsonSchema": bson.M{
			"bsonType": "object",
			"required": []string{"from", "to", "amount"},
			"properties": bson.M{
				"amount": bson.M{
					"bsonType": "double",
					"minimum": 0,
					"description": "Must be greater than 0",
				},
			},
		},
	}
)

type Store struct {
	config	*Config
	db		*mongo.Client
}

func New() (IStore, error) {
	conf, err := loadConfig()
	if err != nil {
		return nil, fmt.Errorf("Failed to load config, err: %v", err)
	}

	db := connectDB(conf)

	if err := Migrate(db.Database(conf.MongoDatabase), accountCollection, validatorAccount); err != nil {
		fmt.Printf("Failed to create collection: %v\n", err)
	}

	if err := Migrate(db.Database(conf.MongoDatabase), transactionCollection, validatorTransaction); err != nil {
		fmt.Printf("Failed to create collection: %v\n", err)
	}

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

func Migrate(db *mongo.Database, collectionName string, validator primitive.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	opt := options.CreateCollection().SetValidator(validator)
	if err := db.CreateCollection(ctx, collectionName, opt); err != nil {
		return err
	}

	return nil
}

func (s Store) Account() IAccountInterface {
	return accountStore{s}
}

func (s Store) Transaction() ITransactionInterface {
	return transactionStore{s}
}
