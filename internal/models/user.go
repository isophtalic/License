package models

import (
	"time"
)

var USER_TABLE string = "users"

type User struct {
	UserID       *string   `json:"userID" gorm:"type:uuid;column:user_id;primary_key"`
	Email        *string   `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password     *string   `json:"password" gorm:"type:varchar(255);not null"`
	Salt         *string   `json:"salt" gorm:"type:varchar(50);not null"`
	Name         *string   `json:"name" gorm:"type:varchar(255);not null"`
	Role         *string   `json:"role" gorm:"type:varchar(10);not null"`
	Status       *bool     `json:"status" gorm:"type:bool;not null"`
	LastLoggedIn time.Time `json:"last_logged_in" gorm:"type:timestamp;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}
