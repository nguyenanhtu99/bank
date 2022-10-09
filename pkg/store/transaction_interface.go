package store

import "bank/pkg/model"

type ITransactionInterface interface {
	Create(account *model.Transaction) error
	GetByUserName(userName string) (*[]model.Transaction, error)
}