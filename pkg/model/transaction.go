package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID 			primitive.ObjectID	`json:"id,omitempty" bson:"_id,omitempty"`
	From 		string				`json:"from"`
	To 			string				`json:"to"`
	Amount		float64				`json:"amount"`
	CreatedAt	int64				`json:"createdAt"`
}