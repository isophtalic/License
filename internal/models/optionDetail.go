package models

import "time"

var OPTION_DETAIL_TABLE string = "option_details"

type OptionDetail struct {
	OptionDetailID     *string    `json:"optionDetailID,omitempty" gorm:"column:option_detail_id;type:uuid;primary_key"`
	Key                *string    `json:"key,omitempty" gorm:"column:key;type:varchar(30);not null"`
	Value              *string    `json:"value,omitempty" gorm:"column:value;type:varchar(30);not null"`
	CreatedAt          *time.Time `json:"createdAt,omitempty" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt          *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at;type:timestamp;not null"`
	ProductOptionID_FK *string    `json:"productOptionID,omitempty" gorm:"column:product_option_id;type:uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// ProductOption      ProductOption `json:"-" gorm:"foreignKey:ProductOptionID_FK;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
