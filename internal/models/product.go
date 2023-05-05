package models

import "time"

var PRODUCT_TABLE string = "products"

type Product struct {
	ProductID    *string    `json:"productID,omitempty" gorm:"column:product_id;type:uuid;primary_key"`
	Name         *string    `json:"name,omitempty" gorm:"column:name;type:varchar(50);not null;unique"`
	Description  *string    `json:"description,omitempty" gorm:"column:description;type:text;not null"`
	Status       *bool      `json:"status,omitempty" gorm:"column:status;type:bool;default:true;not null"`
	Company      *string    `json:"company,omitempty" gorm:"column:company;type:varchar(100);not null"`
	Email        *string    `json:"email,omitempty" gorm:"column:email;type:varchar(50);not null"`
	Phone        *string    `json:"phone,omitempty" gorm:"column:phone;type:varchar(11);not null"`
	Address      *string    `json:"address,omitempty" gorm:"column:address;type:varchar(100);not null"`
	CreatedAt    *time.Time `json:"createdAt,omitempty" gorm:"column:created_at;type:timestamp"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at;type:timestamp"`
	CreatorID_FK *string    `json:"creatorID,omitempty" gorm:"column:creator_id;type:uuid"`

	Creator        *User           `json:"creator,omitempty" gorm:"foreignKey:CreatorID_FK;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductOptions []ProductOption `json:"productOptions" gorm:"foreignKey:ProductID_FK;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Key            *Key            `json:"key" gorm:"foreignKey:ProductID_FK;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
