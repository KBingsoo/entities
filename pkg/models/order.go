package models

import "strings"

type Order struct {
	ID     string   `json:"id" bson:"_id,omitempty"`
	UserID string   `json:"user_id" bson:"user_id"`
	Cards  []string `json:"itens" bson:"itens"`
	Status Status   `json:"status" bson:"status"`
}

type Status string

const (
	PROCESSING        Status = "processing"
	PROCCESINGPAYMENT Status = "processing_payment"
	COMPLETED         Status = "completed"
	CANCELED          Status = "canceled"
)

func NewStatus(status string) Status {
	switch strings.ToLower(status) {
	case "processing":
		return PROCESSING
	case "processing_payment":
		return PROCCESINGPAYMENT
	case "completed":
		return COMPLETED
	case "canceled":
		return CANCELED
	default:
		return PROCESSING
	}
}
