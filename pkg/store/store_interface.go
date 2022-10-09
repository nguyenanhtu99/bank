package store

type IStore interface {
	Account()		IAccountInterface
	Transaction()	ITransactionInterface
}