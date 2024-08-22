package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	Quantity      *string            `json:"number_of_guests" validate:required`
	Unit_price    *float64           `json:"unit_price" validate:requried,eq=S|eq=M|eq=L`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	Food_id       *string            `json:"food_id"`
	Order_item_id string             `json:"order_item_id"`
	Order_id      string             `json:"order_id" validate:required`
}
