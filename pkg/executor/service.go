package executor

import (
	"bank/pkg/store"
	"fmt"
)

type service struct {
	store   store.IStore
}

func New() (IService, error) {
	s, err := store.New()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare store, err: %v", err)
	}

	return &service{
		store:   s,
	}, nil
}