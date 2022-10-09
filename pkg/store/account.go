package store

import (
	"bank/pkg/model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const accountCollection = "accounts"

type accountStore struct {
	Store
}

func (s accountStore) Create(account *model.Account) error {
	fmt.Println("Store: creating account ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := s.db.Database(s.config.MongoDatabase).Collection(accountCollection)
	if _, err := collection.InsertOne(ctx, account); err != nil {
		fmt.Printf("Store: failed to insert account: %v", err)
		return err
	}

	fmt.Println("Store: create account successfully!")

	return nil
}

func (s accountStore) GetByUserName(userName string) (*model.Account, error) {
	fmt.Println("Store: getting account ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var result *model.Account
	defer cancel()

	collection := s.db.Database(s.config.MongoDatabase).Collection(accountCollection)
	if err := collection.FindOne(ctx, bson.M{"username": userName}).Decode(&result); err != nil {
		fmt.Printf("Store: failed to find account: %v\n", err)

		return nil, err
	}

	fmt.Println("Store: get account successfully!")

	return result, nil
}

func (s accountStore) UpdateBalance(account *model.Account) (error) {
	fmt.Printf("Store: updating account: %v ...\n", account.UserName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{"balance": account.Balance}
	collection := s.db.Database(s.config.MongoDatabase).Collection(accountCollection)
	if _, err := collection.UpdateOne(ctx, bson.M{"_id": account.ID}, bson.M{"$set": update}); err != nil {
		fmt.Printf("Store: failed to update account: %v\n", err)
		return err
	}

	fmt.Println("Store: update account successfully!")

	return nil
}