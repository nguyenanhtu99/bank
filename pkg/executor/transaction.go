package executor

import (
	"bank/pkg/model"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
)

type CreateTransactionInput struct {
	From 	string 	`json:"from" validate:"required"`
	To		string 	`json:"to" validate:"required,nefield=From"`
	Amount	float64 `json:"amount" validate:"gt=0"`
}

func (svc *service) CreateTransaction(input *CreateTransactionInput) error {
	fmt.Println("Executor: creating transaction ...")

	var transaction model.Transaction

	if err := copier.Copy(&transaction, input); err != nil {
		fmt.Printf("Executor: failed to copy input to transaction: %v\n", err)
		return err
	}

	transaction.CreatedAt = time.Now().Unix()

	if err := svc.store.Transaction().Create(&transaction); err != nil {
		fmt.Printf("Executor: failed to create transaction: %v\n", err)
		return err
	}

	fmt.Println("Executor: create transaction successfully!")

	return nil
}

func (svc *service) GetTransaction(userName string) (*[]model.Transaction, error) {
	fmt.Println("Executor: getting account ...")

	transaction, err := svc.store.Transaction().GetByUserName(userName)
	if err != nil {
		fmt.Printf("Executor: failed to get account: %v\n", err)
		return nil, err
	}

	fmt.Println("Executor: get account successfully!")

	return transaction, nil
}
