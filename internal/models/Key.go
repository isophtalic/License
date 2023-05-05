package models

import "time"

type Key struct {
	KeyID        *string    `json:"key_id,omitempty" gorm:"column:key_id;type:uuid;primary_key"`
	Type         *string    `json:"type,omitempty" gorm:"column:type;type:varchar(20);not null"`
	PrivateKey   *string    `json:"privateKey,omitempty" gorm:"column:private_key;type:text;not null"`
	PublicKey    *string    `json:"publicKey,omitempty" gorm:"column:public_key;type:text;not null"`
	CreatedAt    *time.Time `json:"createdAt,omitempty" gorm:"column:created_at;type:timestamp"`
	CreatorID_FK *string    `json:"creatorID,omitempty" gorm:"column:creator_id;type:uuid"`
	ProductID_FK *string    `json:"productID,omitempty" gorm:"column:product_id;type:uuid;foreignKey:ProductID_FK;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Creator      *User      `json:"creator" gorm:"foreignKey:CreatorID_FK;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
