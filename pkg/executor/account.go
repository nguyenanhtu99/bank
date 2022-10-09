package executor

import (
	"bank/pkg/model"
	"fmt"

	"github.com/jinzhu/copier"
)

const defaultBalance = 10

type CreateAccountInput struct {
	UserName string `json:"userName" validate:"required,alphanum,len=5"`
}

func (svc *service) CreateAccount(input *CreateAccountInput) error {
	fmt.Println("Executor: creating account ...")

	var account model.Account

	if err := copier.Copy(&account, input); err != nil {
		fmt.Printf("Executor: failed to copy input to account: %v\n", err)
		return err
	}

	account.Balance = defaultBalance

	if err := svc.store.Account().Create(&account); err != nil {
		fmt.Printf("Executor: failed to create account: %v\n", err)
		return err
	}

	fmt.Println("Executor: create account successfully!")

	return nil
}

func (svc *service) GetAccount(userName string) (*model.AccountResponse, error) {
	fmt.Printf("Executor: getting account: %s ...\n", userName)

	account, err := svc.store.Account().GetByUserName(userName)
	if err != nil {
		fmt.Printf("Executor: failed to get account: %v\n", err)
		return nil, err
	}

	response, err := svc.toAccountResponse(account)
	if err != nil {
		fmt.Printf("Executor: failed to get response: %v\n", err)
		return nil, err
	}

	fmt.Println("Executor: get account successfully!")

	return response, nil
}

func (svc *service) toAccountResponse(account *model.Account) (*model.AccountResponse, error) {
	fmt.Printf("Getting response: %s ...\n", account.UserName)

	var result model.AccountResponse

	if err := copier.Copy(&result, account); err != nil {
		fmt.Printf("Failed to copy account to response: %v\n", err)
		return nil, err
	}

	history, _ := svc.store.Transaction().GetByUserName(account.UserName)
	result.History = history

	fmt.Println("Get response successfully!")

	return &result, nil
}