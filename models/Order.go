package models

import (
	"time"
)

type Order struct {
	Id           int       `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	CustomerName string    `gorm:"type:varchar(50);not null" json:"customer_name"`
	Quantity     int       `gorm:"type:integer;not null" json:"quantity"`
	Price        int       `gorm:"type:integer;not null" json:"price"`
	CreateDate   time.Time `gorm:"type:datetime;not null" json:"create_date"`
	ProductId    int
	Product      Product `gorm:"foreignKey:ProductId"`
}

func (Order) TableName() string {
	return "orders"
}
