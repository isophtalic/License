package models

import (
	"time"

	"github.com/lib/pq"
)

type Customer struct {
	CustomerID   *string        `json:"customerID" gorm:"type:uuid;primary_key"`
	Email        *string        `json:"email" gorm:"type:varchar(100);unique;not null"`
	Name         *string        `json:"name" gorm:"type:varchar(50);unique; not null"`
	Company      *string        `json:"company" gorm:"type:varchar(30); not null"`
	Phone_Number *string        `json:"phone_number" gorm:"type:varchar(11);column:phone_number;unique;not null"`
	Organization *string        `json:"organization" gorm:"type:varchar(20);not null"`
	BirthDay     time.Time      `json:"birthDay" gorm:"type:timestamp;not null"`
	Address      *string        `json:"address" gorm:"type: varchar(100); not null"`
	Description  *string        `json:"description" gorm:"type: text; not null"`
	Tags         pq.StringArray `json:"tags" gorm:"type:text[]; not null"`
	Enable       *bool          `json:"enable" gorm:"type:bool; not null"`
	CreatedAt    time.Time      `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"type:timestamp;not null"`
}
