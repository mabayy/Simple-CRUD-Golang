package models

import (
	"time"
)

type Order struct {
	OrderId      uint      `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items"`
}
