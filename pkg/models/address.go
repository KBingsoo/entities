package models

type Address struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	UserID string `json:"user_id" bson:"user_id"`
	Street string `json:"street" bson:"street"`
	Number string `json:"number" bson:"number"`

	Country string `json:"country" bson:"country"`
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`

	Zip string `json:"zip" bson:"zip"`
}
