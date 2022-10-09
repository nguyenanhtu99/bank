package executor

import "bank/pkg/model"

type IService interface {
	CreateAccount(input *CreateAccountInput) error
	GetAccount(userName string) (*model.AccountResponse, error)

	CreateTransaction(input *CreateTransactionInput) error
	GetTransaction(userName string) (*[]model.Transaction, error)
}
