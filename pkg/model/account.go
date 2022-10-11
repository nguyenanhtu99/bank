package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID			primitive.ObjectID	`json:"id,omitempty" bson:"_id,omitempty"`
	UserName	string				`json:"userName"`
	Balance		float64				`json:"balance"`
}

type AccountResponse struct {
	ID 			primitive.ObjectID	`json:"id,omitempty" bson:"_id,omitempty"`
	UserName	string				`json:"userName"`
	Balance		float64				`json:"balance"`
	History		*[]Transaction		`json:"history"`
}