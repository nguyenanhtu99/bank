package store

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	accountValidator = bson.M{
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
	transationValidator = bson.M{
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

	if err := accountMigrate(db.Database(conf.MongoDatabase)); err != nil {
		fmt.Printf("Failed to create collection: %v\n", err)
	}

	if err := transactionMigrate(db.Database(conf.MongoDatabase)); err != nil {
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
 
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	err = client.Connect(ctx)
	defer cancel()
 
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to mongoDB")
	return client
}

func accountMigrate(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	optValidator := options.CreateCollection().SetValidator(accountValidator)
	if err := db.CreateCollection(ctx, accountCollection, optValidator); err != nil {
		return err
	}

	model := mongo.IndexModel{
		Keys: bson.M{"username": "text"}, 
		Options: options.Index().SetUnique(true),
	}
	if _, err := db.Collection(accountCollection).Indexes().CreateOne(ctx, model); err != nil {
		return err
	}

	return nil
}

func transactionMigrate(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	optValidator := options.CreateCollection().SetValidator(transationValidator)
	if err := db.CreateCollection(ctx, transactionCollection, optValidator); err != nil {
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
