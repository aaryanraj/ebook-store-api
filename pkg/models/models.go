package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	FullName  string    `gorm:"size:255;not null;" json:"full_name"`
	UserName  string    `gorm:"size:255;not null;unique" json:"nick_name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
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
