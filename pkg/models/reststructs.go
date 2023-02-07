package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"size:255;not null;" json:"full_name"`
	UserName string `gorm:"size:255;not null;unique" json:"nick_name"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
}

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Publication string `json:"publication" binding:"required"`
}

type UserResponse struct {
	Status_Code int    `json:"statusCode"`
	Message     string `json:"message"`
	More_info   string `json:"more_info"`
}
