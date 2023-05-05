package dto

import "time"

type OptionDetailDTO struct {
	Key                *string    `json:"key,omitempty"`
	Value              *string    `json:"value,omitempty"`
	CreatedAt          *time.Time `json:"createdAt,omitempty"`
	UpdatedAt          *time.Time `json:"updatedAt,omitempty"`
	ProductOptionID_FK *string    `json:"productOptionID,omitempty"`
}
