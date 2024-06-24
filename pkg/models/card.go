package models

type Card struct {
	ID    string  `json:"id" bson:"_id,omitempty"`
	Name  string  `json:"name" bson:"name"`
	Slug  string  `json:"slug" bson:"slug"`
	Qty   int     `json:"qty" bson:"qty,omitempty"`
	Price float32 `json:"price" bson:"price"`
}
