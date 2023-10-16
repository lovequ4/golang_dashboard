package models

import (
	"time"
)

type Product struct {
	Id          int       `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	ProductName string    `gorm:"type:varchar(50);unique;not null" json:"product_name"`
	Quantity    int       `gorm:"type:integer;not null" json:"quantity"`
	Price       int       `gorm:"type:integer;not null" json:"price"`
	CreateDate  time.Time `gorm:"type:datetime;not null" json:"create_date"`
}

func (Product) TableName() string {
	return "products"
}
