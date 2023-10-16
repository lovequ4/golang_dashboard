package models

import "time"

type User struct {
	Id       int       `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Name     string    `gorm:"type:varchar(25);not null" json:"name"`
	Password string    `gorm:"type:varchar(255);not null" json:"password"`
	Email    string    `gorm:"type:varchar(50);unique;not null" json:"email"`
	Role     string    `gorm:"type:varchar(10);not null;default:'employee'" json:"role"`
	Date     time.Time `gorm:"type:datetime;not null" json:"date"`
}

func (User) TableName() string {
	return "users"
}
