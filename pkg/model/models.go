package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
}

type Order struct {
	gorm.Model
	User      User    `gorm:"foreignKey:UserID"`
	Product   Product `gorm:"foreignKey:ProductID"`
	UserID    uint
	ProductID uint
	Quantity  int `json:"quantity"`
}
