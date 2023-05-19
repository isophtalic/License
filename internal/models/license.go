package models

import (
	"time"

	"github.com/lib/pq"
)

type License struct {
	LicenseID   *string        `json:"license_id" gorm:"column:license_id;type:uuid;primary_key"`
	Name        *string        `json:"name" gorm:"column:name;type:varchar(50);unique;not null"`
	Option      *string        `json:"option" gorm:"column:option;type:varchar(10);not null"`
	ExpiryDate  time.Time      `json:"expiry_date" gorm:"column:expiry_date;type:timestamp;not null"`
	Description *string        `json:"description" gorm:"column:description;type:text;not null"`
	Tag         pq.StringArray `json:"tags" gorm:"column:tags;type:text[];not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null"`
	CustomerID  *string        `json:"customerID" gorm:"foreignKey:CustomerID_FK; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductID   *string        `json:"productID" gorm:"foreignKey:ProductID_FK; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
