package models

import "time"

var PRODUCT_OPTION_TABLE string = "product_options"

type ProductOption struct {
	OptionID     *string    `json:"optionID,omitempty" gorm:"column:option_id;type:uuid;primary_key"`
	Name         *string    `json:"name,omitempty" gorm:"column:name;type:varchar(50);not null"`
	Description  *string    `json:"description,omitempty" gorm:"column:description;type:text;not null"`
	Enable       *bool      `json:"enable,omitempty" gorm:"column:enable;type:bool;default:true;not null"`
	CreatedAt    *time.Time `json:"createdAt,omitempty" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at;type:timestamp;not null"`
	CreatorID_FK *string    `json:"creatorID,omitempty" gorm:"column:creator_id;type:uuid;"`
	ProductID_FK *string    `json:"productID,omitempty" gorm:"column:product_id;type:uuid"`

	Creator       *User           `json:"creator" gorm:"foreignKey:CreatorID_FK;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Product       *Product        `json:"-" gorm:"foreignKey:ProductID_FK;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	OptionDetails *[]OptionDetail `json:"optionDetails" gorm:"foreignKey:ProductOptionID_FK;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
