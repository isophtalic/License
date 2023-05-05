package dto

import "git.cyradar.com/license-manager/backend/internal/models"

type ProductOptionDTO struct {
	Name         *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	Enable       *bool   `json:"enable,omitempty"`
	CreatorID_FK *string `json:"creatorID,omitempty"`
	ProductID_FK *string `json:"productID,omitempty"`
	// OptionDetails []OptionDetailDTO `json:"optionDetails,omitempty"`
}

func ToProductOption(pdDTO *ProductOptionDTO) *models.ProductOption {
	return &models.ProductOption{
		Name:         pdDTO.Name,
		Description:  pdDTO.Description,
		Enable:       pdDTO.Enable,
		CreatorID_FK: pdDTO.CreatorID_FK,
		ProductID_FK: pdDTO.ProductID_FK,
	}
}
