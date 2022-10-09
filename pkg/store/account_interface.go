package store

import (
	"bank/pkg/model"
)

type IAccountInterface interface {
	Create(account *model.Account) error
	GetByUserName(userName string) (*model.Account, error)
	UpdateBalance(account *model.Account) error
}