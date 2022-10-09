package executor

import (
	"bank/pkg/model"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
)

type CreateTransactionInput struct {
	From string `json:"from" validate:"required"`
	To string `json:"to" validate:"required"`
	Amount float64 `json:"amount" validate:"gt=0"`
}

func (svc *service) CreateTransaction(input *CreateTransactionInput) error {
	fmt.Println("Executor: creating transaction ...")

	var transaction model.Transaction

	if err := copier.Copy(&transaction, input); err != nil {
		fmt.Printf("Executor: failed to copy input to transaction: %v\n", err)
		return err
	}

	transaction.CreatedAt = time.Now().Unix()

	if err := svc.TransferAvailable(input); err != nil {
		fmt.Printf("Executor: Transfer not available: %v\n", err)
		return err
	}

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

func (svc *service) TransferAvailable(input *CreateTransactionInput) error {
	if input.From == input.To {
		return fmt.Errorf("Sender and receiver must be different")
	}

	sender, err := svc.store.Account().GetByUserName(input.From)
	if err != nil {
		return fmt.Errorf("Sender not found")
	}

	if sender.Balance < input.Amount {
		return fmt.Errorf("Insufficient funds")
	}

	receiver, err := svc.store.Account().GetByUserName(input.To)
	if err != nil {
		return fmt.Errorf("Receiver not found")
	}

	sender.Balance -= input.Amount
	receiver.Balance += input.Amount

	if err := svc.store.Account().UpdateBalance(sender); err != nil {
		return fmt.Errorf("Failed to update balance: %v", err)
	}

	if err := svc.store.Account().UpdateBalance(receiver); err != nil {
		return fmt.Errorf("Failed to update balance: %v", err)
	}

	return nil
}