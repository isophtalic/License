package models

import "time"

type LicenseConfig struct {
	ConfigID  *string   `json:"config_id" gorm:"column:config_id;type:uuid;primary_key"`
	Name      *string   `json:"name,omitempty" gorm:"column:name;type:varchar(30);unique;not null"`
	Key       *string   `json:"key,omitempty" gorm:"column:key;type:varchar(30);not null"`
	Value     *string   `json:"value,omitempty" gorm:"column:value;type:varchar(30);not null"`
	Status    *bool     `json:"status,omitempty" gorm:"column:status;type:varchar(30);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null"`
	LicenseID *string   `json:"license_id" gorm:"foreignKey:LicenseID_FK; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
