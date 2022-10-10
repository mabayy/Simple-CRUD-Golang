package models

import "gorm.io/gorm"

type Item struct {
	OrderId     uint   `gorm:"foreignKey:order_id" json:"order_id"`
	ItemId      uint   `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	return
}
