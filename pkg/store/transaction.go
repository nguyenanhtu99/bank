package store

import (
	"bank/pkg/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

const transactionCollection = "transactions"

type transactionStore struct {
	Store
}

func (s transactionStore) Create(transaction *model.Transaction) error {
	fmt.Println("Store: creating transaction ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accountCollection := s.db.Database(s.config.MongoDatabase).Collection(accountCollection)
	transactionCollection := s.db.Database(s.config.MongoDatabase).Collection(transactionCollection)

	wc := writeconcern.New(writeconcern.WMajority())
	txnOpts := options.Transaction().SetWriteConcern(wc)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		updateSender := bson.M{
			"$inc": bson.M{
				"balance": transaction.Amount * -1,
			},
		}
		if err := accountCollection.FindOneAndUpdate(sessCtx, bson.D{{Key: "username", Value: transaction.From}}, updateSender).Err(); err != nil {
			fmt.Printf("Store: failed to update sender: %v\n", err)
			return nil, err
		}

		fmt.Println("Store: update sender successfully!")

		updateReceiver := bson.M{
			"$inc": bson.M{
				"balance": transaction.Amount,
			},
		}
		if err := accountCollection.FindOneAndUpdate(sessCtx, bson.D{{Key: "username", Value: transaction.To}}, updateReceiver).Err(); err != nil {
			fmt.Printf("Store: failed to update receiver: %v\n", err)
			return nil, err
		}

		fmt.Println("Store: update receiver successfully!")

		if _, err := transactionCollection.InsertOne(ctx, transaction); err != nil {
			fmt.Printf("Store: failed to insert transaction: %v", err)
			return nil, err
		}

		fmt.Println("Store: create transaction successfully!")

		return nil, nil
	}

	session, err := s.db.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	if _, err := session.WithTransaction(ctx, callback, txnOpts); err != nil {
		return err
	}

	fmt.Println("Store: create transaction successfully!")

	return nil
}

func (s transactionStore) GetByUserName(userName string) (*[]model.Transaction, error) {
	fmt.Println("Store: getting transaction ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var result []model.Transaction
	defer cancel()

	collection := s.db.Database(s.config.MongoDatabase).Collection(transactionCollection)
	filter := bson.D{
		{Key: "$or",
			Value: bson.A{
				bson.D{{Key: "from", Value: userName}},
				bson.D{{Key: "to", Value: userName}},
			}},
	}
	sort := bson.D{{Key: "createdat", Value: -1}}
	opts := options.Find().SetSort(sort)
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		fmt.Printf("Store: failed to find transaction: %v\n", err)

		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var transaction model.Transaction
		if err = cursor.Decode(&transaction); err != nil {
			fmt.Printf("Store: failed to decode transaction: %v\n", err)

			return nil, err
		}

		result = append(result, transaction)
	}

	fmt.Println("Store: get transaction successfully!")

	return &result, nil
}
